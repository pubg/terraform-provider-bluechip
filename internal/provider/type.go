package provider

import (
	"context"

	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwflex"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ fwtype.TypeHelper[ProviderAuths] = &ProviderAuthType{}

type ProviderAuths struct {
	Items []ProviderAuth
}

type ProviderAuth interface {
	_ProviderAuth()
}

type TokenAuth struct {
	ProviderAuth

	Token string
}

type BasicAuth struct {
	ProviderAuth

	Username string
	Password string
}

type AwsAuth struct {
	ProviderAuth

	ClusterName     string
	AccessKey       string
	SecretAccessKey string
	SessionToken    string
	Region          string
	Profile         string
}

type OidcAuth struct {
	ProviderAuth

	ValidatorName string
	Token         *string
	TokenEnv      *string
}

type ProviderAuthType struct {
}

func (t ProviderAuthType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"token": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    0,
			MaxItems:    1,
			Description: "Only one of token, basic, aws, oidc can be specified.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"token": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_TOKEN", ""),
				},
			}},
		},
		"basic": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    0,
			MaxItems:    1,
			Description: "Only one of token, basic, aws, oidc can be specified.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
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
			}},
		},
		"aws": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    0,
			MaxItems:    1,
			Description: "Only one of token, basic, aws, oidc can be specified.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"cluster_name": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_AWS_CLUSTER_NAME", ""),
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
					Optional: true,
				},
				"profile": {
					Type:     schema.TypeString,
					Optional: true,
				},
			}},
		},
		"oidc": {
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    0,
			MaxItems:    1,
			Description: "Only one of token, basic, aws, oidc can be specified.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"validator_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("BLUECHIP_OIDC_TOKEN", ""),
				},
				"token_env": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "",
				},
			}},
		},
	}

	blockSchema := &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MinItems: 0,
		Elem:     &schema.Resource{Schema: innerSchema},
	}
	return blockSchema
}

func (t ProviderAuthType) Expand(ctx context.Context, d *schema.ResourceData, out *ProviderAuths) diag.Diagnostics {
	authFlowAttrs := fwflex.ExpandMapList(d.Get("auth_flow").([]any))
	for _, authFlowAttr := range authFlowAttrs {
		var count = 0
		if rawAttr, ok := authFlowAttr["token"]; ok && len(rawAttr.([]any)) > 0 {
			attr := fwflex.ExpandMapList(rawAttr.([]any))[0]
			out.Items = append(out.Items, &TokenAuth{
				Token: attr["token"].(string),
			})
			count++
		}
		if rawAttr, ok := authFlowAttr["basic"]; ok && len(rawAttr.([]any)) > 0 {
			attr := fwflex.ExpandMapList(rawAttr.([]any))[0]
			out.Items = append(out.Items, &BasicAuth{
				Username: attr["username"].(string),
				Password: attr["password"].(string),
			})
			count++
		}
		if rawAttr, ok := authFlowAttr["aws"]; ok && len(rawAttr.([]any)) > 0 {
			attr := fwflex.ExpandMapList(rawAttr.([]any))[0]
			out.Items = append(out.Items, &AwsAuth{
				ClusterName:     fwflex.ExpandStringFromMap(attr, "cluster_name"),
				AccessKey:       fwflex.ExpandStringFromMap(attr, "access_key"),
				SecretAccessKey: fwflex.ExpandStringFromMap(attr, "secret_access_key"),
				SessionToken:    fwflex.ExpandStringFromMap(attr, "session_token"),
				Region:          fwflex.ExpandStringFromMap(attr, "region"),
				Profile:         fwflex.ExpandStringFromMap(attr, "profile"),
			})
			count++
		}
		if rawAttr, ok := authFlowAttr["oidc"]; ok && len(rawAttr.([]any)) > 0 {
			attr := fwflex.ExpandMapList(rawAttr.([]any))[0]
			out.Items = append(out.Items, &OidcAuth{
				ValidatorName: attr["validator_name"].(string),
				Token:         fwflex.ExpandStringPFromMap(attr, "token"),
				TokenEnv:      fwflex.ExpandStringPFromMap(attr, "token_env"),
			})
			count++
		}

		if count >= 2 || count == 0 {
			return diag.Errorf("Only one of token, basic, aws, oidc can be specified. There are %d of auth blocks specified", count)
		}
	}
	return nil
}

func (t ProviderAuthType) Flatten(in ProviderAuths) map[string]any {
	return nil
}
