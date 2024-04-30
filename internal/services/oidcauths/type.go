package oidcauths

import (
	"context"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwflex"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
					"from_path_resolver": {
						Type:     schema.TypeString,
						Optional: !t.Computed,
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

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip.OidcAuthSpec) diag.Diagnostics {
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
		for _, rawAttributeMapping := range fwflex.ExpandMapList(attr["attribute_mapping"].([]any)) {
			mapping := bluechip.AttributeMapping{
				From: rawAttributeMapping["from"].(string),
				To:   rawAttributeMapping["to"].(string),
			}
			if rawAttributeMapping["from_path_resolver"] != nil {
				mapping.FromPathResolver = rawAttributeMapping["from_path_resolver"].(string)
			}
			out.AttributeMapping = append(out.AttributeMapping, mapping)
		}
	}
	return nil
}

func (t SpecType) Flatten(in bluechip.OidcAuthSpec) map[string]any {
	attr := map[string]any{
		"username_claim":  in.UsernameClaim,
		"username_prefix": in.UsernamePrefix,
		"issuer":          in.Issuer,
		"client_id":       in.ClientId,
		"required_claims": in.RequiredClaims,
		"groups_claim":    in.GroupsClaim,
		"groups_prefix":   in.GroupsPrefix,
	}
	if len(in.AttributeMapping) > 0 {
		var attributeMapping []map[string]any
		for _, mapping := range in.AttributeMapping {
			attributeMapping = append(attributeMapping, map[string]any{
				"from":               mapping.From,
				"from_path_resolver": mapping.FromPathResolver,
				"to":                 mapping.To,
			})
		}
		attr["attribute_mapping"] = attributeMapping
	}
	return attr
}
