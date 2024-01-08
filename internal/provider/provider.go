package provider

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_authenticator"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwlog"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc("BLUECHIP_ADDRESS", ""),
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_TOKEN", ""),
			},
			"basic_auth": {
				Type:     schema.TypeList,
				MinItems: 0,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_USERNAME", ""),
						},
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_PASSWORD", ""),
						},
					},
				},
			},
			"aws_auth": {
				Type:     schema.TypeList,
				MinItems: 0,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_CLUSTER_NAME", ""),
						},
						"access_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret_access_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_token": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Required: true,
						},
						"profile": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"oidc_auth": {
				Type:     schema.TypeList,
				MinItems: 0,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"validator_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"token": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_OIDC_TOKEN", ""),
						},
					},
				},
			},
		},
		ResourcesMap:   resourceMap,
		DataSourcesMap: dataSourceMap,
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return providerConfigure(ctx, d, p.TerraformVersion)
	}

	return p
}

type ProviderModel struct {
	Version string
	Client  *bluechip_client.Client
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, providerVersion string) (interface{}, diag.Diagnostics) {
	config := initializeConfiguration(d)
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
	Address   string
	Token     *string
	BasicAuth *struct {
		Username string
		Password string
	}
	AwsAuth *struct {
		ClusterName     string
		AccessKey       *string
		SecretAccessKey *string
		SessionToken    *string
		Region          string
		Profile         *string
	}
	OidcAuth *struct {
		ValidatorName string
		Token         string
	}
}

func initializeConfiguration(d *schema.ResourceData) ProviderConfig {
	var config ProviderConfig
	config.Address = d.Get("address").(string)

	if v, ok := d.GetOk("token"); ok {
		config.Token = v.(*string)
	}

	// Copilot wrote this code
	if v, ok := d.GetOk("basic_auth"); ok {
		config.BasicAuth = &struct {
			Username string
			Password string
		}{
			Username: v.([]interface{})[0].(map[string]interface{})["username"].(string),
			Password: v.([]interface{})[0].(map[string]interface{})["password"].(string),
		}
	}

	if v, ok := d.GetOk("aws_auth"); ok {
		config.AwsAuth = &struct {
			ClusterName     string
			AccessKey       *string
			SecretAccessKey *string
			SessionToken    *string
			Region          string
			Profile         *string
		}{
			ClusterName:     v.([]interface{})[0].(map[string]interface{})["cluster_name"].(string),
			AccessKey:       v.([]interface{})[0].(map[string]interface{})["access_key"].(*string),
			SecretAccessKey: v.([]interface{})[0].(map[string]interface{})["secret_access_key"].(*string),
			SessionToken:    v.([]interface{})[0].(map[string]interface{})["session_token"].(*string),
			Region:          v.([]interface{})[0].(map[string]interface{})["region"].(string),
			Profile:         v.([]interface{})[0].(map[string]interface{})["profile"].(*string),
		}
	}

	if v, ok := d.GetOk("oidc_auth"); ok {
		config.OidcAuth = &struct {
			ValidatorName string
			Token         string
		}{
			ValidatorName: v.([]interface{})[0].(map[string]interface{})["validator_name"].(string),
			Token:         v.([]interface{})[0].(map[string]interface{})["token"].(string),
		}
	}

	return config
}

func initializeBluechipToken(ctx context.Context, authClient *bluechip_authenticator.Client, config ProviderConfig) (string, diag.Diagnostics) {
	if config.Token != nil {
		return *config.Token, nil
	} else if config.BasicAuth != nil {
		token, err := authClient.LoginWithBasic(context.Background(), config.BasicAuth.Username, config.BasicAuth.Password)
		if err != nil {
			return "", diag.FromErr(err)
		}
		return token, nil
	} else if config.AwsAuth != nil {
		token, err := authClient.LoginWithAws(ctx, config.AwsAuth.ClusterName, defaultString(config.AwsAuth.AccessKey, ""), defaultString(config.AwsAuth.SecretAccessKey, ""), defaultString(config.AwsAuth.SessionToken, ""), config.AwsAuth.Region, defaultString(config.AwsAuth.Profile, ""))
		if err != nil {
			return "", diag.FromErr(err)
		}
		return token, nil
	} else if config.OidcAuth != nil {
		token, err := authClient.LoginWithOidc(context.Background(), config.OidcAuth.Token, config.OidcAuth.ValidatorName)
		if err != nil {
			return "", diag.FromErr(err)
		}
		return token, nil
	}
	return "", diag.Errorf("either token, basic_auth, aws_auth or oidc_auth must be specified")
}

func defaultString(s *string, def string) string {
	if s == nil {
		return def
	}
	return *s
}

var resourceMap = map[string]*schema.Resource{}

func RegisterResource(name string, resource *schema.Resource) {
	resourceMap[name] = resource
}

var dataSourceMap = map[string]*schema.Resource{}

func RegisterDataSource(name string, dataSource *schema.Resource) {
	dataSourceMap[name] = dataSource
}
