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
	Computed bool
}

func (t SpecType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"username_claim": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"username_prefix": {
			Type:     schema.TypeString,
			Optional: !t.Computed,
			Computed: t.Computed,
		},
		"issuer": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"required_claims": {
			Type:     schema.TypeSet,
			Optional: !t.Computed,
			Computed: t.Computed,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"groups_claim": {
			Type:     schema.TypeString,
			Optional: !t.Computed,
			Computed: t.Computed,
		},
		"groups_prefix": {
			Type:     schema.TypeString,
			Optional: !t.Computed,
			Computed: t.Computed,
		},
		"attribute_mapping": {
			Type:     schema.TypeList,
			Optional: !t.Computed,
			Computed: t.Computed,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"from": {
						Type:     schema.TypeString,
						Required: !t.Computed,
						Computed: t.Computed,
					},
					"to": {
						Type:     schema.TypeString,
						Required: !t.Computed,
						Computed: t.Computed,
					},
				},
			},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.OidcAuthSpec) diag.Diagnostics {
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
	if attr["groups_claim"] != nil {
		out.GroupsClaim = fwtype.String(attr["groups_claim"].(string))
	}
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

func (t SpecType) Flatten(in bluechip_models.OidcAuthSpec) map[string]any {
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

	return attr
}
