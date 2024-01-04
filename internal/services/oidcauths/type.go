package oidcauths

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type SpecType struct {
}

func (SpecType) Schema(computed bool) *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"username_claim": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"username_prefix": {
			Type:     schema.TypeString,
			Optional: !computed,
			Computed: computed,
		},
		"issuer": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"required_claims": {
			Type:     schema.TypeSet,
			Optional: !computed,
			Computed: computed,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"groups_claim": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"groups_prefix": {
			Type:     schema.TypeString,
			Optional: !computed,
			Computed: computed,
		},
		"attribute_mapping": {
			Type:     schema.TypeList,
			Optional: !computed,
			Computed: computed,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"from": {
						Type:     schema.TypeString,
						Required: !computed,
						Computed: computed,
					},
					"to": {
						Type:     schema.TypeString,
						Required: !computed,
						Computed: computed,
					},
				},
			},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.OidcAuthSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.UsernameClaim = attr["username_claim"].(string)
	if attr["username_prefix"] != nil {
		out.UsernamePrefix = fwtype.String(attr["username_prefix"].(string))
	}
	out.Issuer = attr["issuer"].(string)
	out.ClientId = attr["client_id"].(string)
	if attr["required_claims"] != nil {
		out.RequiredClaims = fwflex.ExpandStringSet(attr["required_claims"].(*schema.Set))
	}
	out.GroupsClaim = attr["groups_claim"].(string)
	if attr["groups_prefix"] != nil {
		out.GroupsPrefix = fwtype.String(attr["groups_prefix"].(string))
	}
	if attr["attribute_mapping"] != nil {
		rawAttributeMappings := fwflex.ExpandMapList(attr["attribute_mapping"].([]any))
		for _, rawAttributeMapping := range rawAttributeMappings {
			out.AttributeMapping = append(out.AttributeMapping, bluechip_models.AttributeMapping{
				From: rawAttributeMapping["from"].(string),
				To:   rawAttributeMapping["to"].(string),
			})
		}
	}
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.OidcAuthSpec) diag.Diagnostics {
	attr := map[string]any{
		"username_claim":    in.UsernameClaim,
		"username_prefix":   in.UsernamePrefix,
		"issuer":            in.Issuer,
		"client_id":         in.ClientId,
		"required_claims":   in.RequiredClaims,
		"groups_claim":      in.GroupsClaim,
		"groups_prefix":     in.GroupsPrefix,
		"attribute_mapping": in.AttributeMapping,
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
