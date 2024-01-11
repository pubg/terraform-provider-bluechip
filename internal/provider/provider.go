package provider

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_authenticator"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwdiag"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwlog"
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
	Client  *bluechip_client.Client
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, providerVersion string) (interface{}, diag.Diagnostics) {
	config, diags := expandProviderConfig(ctx, d)
	if diags.HasError() {
		return nil, diags
	}
	tflog.Debug(ctx, "Read Provider Config", fwlog.Field("config", config))

	// Validate the auth configuration values
	baseClient := &http.Client{Transport: logging.NewLoggingHTTPTransport(http.DefaultTransport)}
	authClient := bluechip_authenticator.NewClient(baseClient, providerVersion, config.Address)
	token, diags := initializeBluechipToken(ctx, authClient, config)
	if diags.HasError() {
		tflog.Info(ctx, "Failed to Initialize BlueChip Token", fwlog.Field("diags", diags))
		return nil, diags
	}

	client := bluechip_client.NewClient(baseClient, token, providerVersion, config.Address, 3*time.Second, 30*time.Second)
	// WhoAmI API로 사용자 정보 검증
	if _, err := client.Whoami(ctx); err != nil {
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

func initializeBluechipToken(ctx context.Context, authClient *bluechip_authenticator.Client, config *ProviderConfig) (string, diag.Diagnostics) {
	var diags diag.Diagnostics

	for index, authConfig := range config.Auths.Items {
		if tokenAuth, ok := authConfig.(*TokenAuth); ok {
			if tokenAuth.Token != "" {
				return tokenAuth.Token, nil
			}
		} else if basicAuth, ok := authConfig.(*BasicAuth); ok {
			token, err := authClient.LoginWithBasic(context.Background(), basicAuth.Username, basicAuth.Password)
			if err != nil {
				diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.basic", index)...)
				continue
			}
			return token, nil
		} else if awsAuth, ok := authConfig.(*AwsAuth); ok {
			var clusterName string
			if awsAuth.ClusterName == "" {
				awsConfig, err := authClient.GetAwsConfiguration(ctx)
				if err != nil {
					diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.aws", index)...)
					continue
				}
				clusterName = awsConfig.ValidClusterNames[0]
			} else {
				clusterName = awsAuth.ClusterName
			}

			awsOptions := &bluechip_authenticator.AwsOptions{
				ClusterName:     clusterName,
				AccessKey:       awsAuth.AccessKey,
				SecretAccessKey: awsAuth.SecretAccessKey,
				SessionToken:    awsAuth.SessionToken,
				Region:          awsAuth.Region,
				Profile:         awsAuth.Profile,
			}
			token, err := authClient.LoginWithAws(ctx, awsOptions)
			if err != nil {
				diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.aws", index)...)
				continue
			}
			return token, nil
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

			token, err := authClient.LoginWithOidc(ctx, oidcToken, oidcAuth.ValidatorName)
			if err != nil {
				diags = append(diags, fwdiag.FromErrF(err, "Login Failed, path: auth_flow.%d.oidc", index)...)
				continue
			}
			return token, nil
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
