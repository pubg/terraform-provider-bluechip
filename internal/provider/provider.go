package provider

import (
	"context"
	"fmt"
	"os"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_client"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_config"
	"git.projectbro.com/Devops/arcane-client-go/pkg/kube_plugins"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwdiag"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwlog"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"k8s.io/client-go/rest"
)

var providerAuthTyp = &ProviderAuthType{}

func Provider(version string, commit string) func() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("BLUECHIP_ADDR", ""),
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"auth_flow": providerAuthTyp.Schema(),
		},
		ResourcesMap:   resourceMap,
		DataSourcesMap: dataSourceMap,
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return providerConfigure(ctx, d, fmt.Sprintf("%s+%s", version, commit))
	}

	return func() *schema.Provider {
		return p
	}
}

type ProviderModel struct {
	Version string
	Client  *bluechip.BluechipClient
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, providerVersion string) (interface{}, diag.Diagnostics) {
	config, diags := expandProviderConfig(ctx, d)
	if diags.HasError() {
		return nil, diags
	}
	tflog.Debug(ctx, "Read Provider Config", fwlog.Field("config", config))

	// Validate the auth configuration values
	arcaneNonAuthClient, err := initializeArcaneClient(ctx, config.Address, nil)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to initialize the provider's auth client: %v", err))
	}
	token, diags := initializeBluechipToken(ctx, arcaneNonAuthClient, config)
	if diags.HasError() {
		tflog.Info(ctx, "Failed to Initialize BlueChip Token", fwlog.Field("diags", diags))
		return nil, diags
	}

	arcaneAuthenticatedClient, err := initializeArcaneClient(ctx, config.Address, &token)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to initialize the provider's client: %v", err))
	}
	client := bluechip.NewBluechipClient(arcaneAuthenticatedClient)
	// WhoAmI API로 사용자 정보 검증
	if _, err := client.Auths().Whoami(ctx); err != nil {
		tflog.Info(ctx, "Failed to validate the provider's auth configuration values", fwlog.Field("error", err))
		return nil, diag.FromErr(fmt.Errorf("failed to validate the provider's auth configuration values: %v", err))
	}

	model := &ProviderModel{
		Version: providerVersion,
		Client:  client,
	}
	return model, nil
}

type ProviderConfig struct {
	Address string
	Auths   ProviderAuths
}

func expandProviderConfig(ctx context.Context, d *schema.ResourceData) (*ProviderConfig, diag.Diagnostics) {
	var config ProviderConfig
	config.Address = d.Get("address").(string)

	if diags := providerAuthTyp.Expand(ctx, d, &config.Auths); diags.HasError() {
		return nil, diags
	}

	return &config, nil
}

func initializeArcaneClient(ctx context.Context, server string, token *string) (*api_client.ArcaneRestClient, diag.Diagnostics) {
	authRestConfig, err := api_config.SimpleConfig(server)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to parse the provider's auth configuration values: %v", err))
	}
	if token != nil {
		authRestConfig.BearerToken = *token
	}

	restClient, err := rest.RESTClientFor(authRestConfig)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to create REST client: %v", err))
	}
	return api_client.NewArcaneRestClient(restClient), nil
}

func initializeBluechipToken(ctx context.Context, nonAuthClient *api_client.ArcaneRestClient, config *ProviderConfig) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	bluechipClient := bluechip.NewBluechipClient(nonAuthClient)

	for index, authConfig := range config.Auths.Items {
		if tokenAuth, ok := authConfig.(*TokenAuth); ok {
			if tokenAuth.Token != "" {
				return tokenAuth.Token, nil
			}
		} else if basicAuth, ok := authConfig.(*BasicAuth); ok {
			lr, err := kube_plugins.LoginByBasicAuth(ctx, nonAuthClient, basicAuth.Username, basicAuth.Password)
			if err != nil {
				diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.basic", index)...)
				continue
			}
			return lr.Token, nil
		} else if awsAuth, ok := authConfig.(*AwsAuth); ok {
			var clusterName string
			if awsAuth.ClusterName == "" {
				bluechipAwsConfig, err := bluechipClient.Auths().GetAwsConfiguration(ctx)
				if err != nil {
					diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.aws", index)...)
					continue
				}
				clusterName = bluechipAwsConfig.ValidClusterNames[0]
			} else {
				clusterName = awsAuth.ClusterName
			}

			var configLoadOptions []func(*awsConfig.LoadOptions) error
			if awsAuth.AccessKey != "" && awsAuth.SecretAccessKey != "" {
				configLoadOptions = append(configLoadOptions, awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAuth.AccessKey, awsAuth.SecretAccessKey, awsAuth.SessionToken)))
				tflog.Debug(ctx, "Using static credentials from provider parameters")
			}
			if awsAuth.Region != "" {
				configLoadOptions = append(configLoadOptions, awsConfig.WithRegion(awsAuth.Region))
				tflog.Debug(ctx, "Using region from provider parameters")
			}
			if awsAuth.Profile != "" {
				configLoadOptions = append(configLoadOptions, awsConfig.WithSharedConfigProfile(awsAuth.Profile))
				tflog.Debug(ctx, "Using profile from provider parameters")
			}

			cfg, err := awsConfig.LoadDefaultConfig(ctx, configLoadOptions...)
			if err != nil {
				return "", diag.FromErr(err)
			}
			if cfg.Region == "" {
				return "", diag.Errorf("cannot determine region from provider parameters or aws config")
			}

			lr, err := kube_plugins.LoginByAwsConfig(ctx, nonAuthClient, clusterName, cfg)
			if err != nil {
				diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.aws", index)...)
				continue
			}
			return lr.Token, nil
		} else if oidcAuth, ok := authConfig.(*OidcAuth); ok {
			var oidcToken string
			if oidcAuth.TokenEnv != nil {
				oidcToken = os.Getenv(*oidcAuth.TokenEnv)
			}
			if oidcAuth.Token != nil {
				oidcToken = *oidcAuth.Token
			}
			if oidcToken == "" {
				diags = append(diags, diag.Errorf("Failed to resolve oidc token, tokenEnv: %+v, token: %+v, path: auth_flow.%d.oidc", oidcAuth.TokenEnv, oidcAuth.Token, index)...)
				continue
			}

			lr, err := kube_plugins.LoginByOidcAuth(ctx, nonAuthClient, oidcAuth.ValidatorName, oidcToken)
			if err != nil {
				diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.oidc", index)...)
				continue
			}
			return lr.Token, nil
		}
	}
	if len(diags) > 0 {
		return "", diags
	} else {
		return "", diag.Errorf("either token, basic, aws or oidc must be specified")
	}
}

var resourceMap = map[string]*schema.Resource{}

func RegisterResource(name string, resource *schema.Resource) {
	resourceMap[name] = resource
}

var dataSourceMap = map[string]*schema.Resource{}

func RegisterDataSource(name string, dataSource *schema.Resource) {
	dataSourceMap[name] = dataSource
}
