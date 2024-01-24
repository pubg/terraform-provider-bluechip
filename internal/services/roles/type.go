package roles

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
		"statements": {
			Type:     schema.TypeList,
			Required: !t.Computed,
			Computed: t.Computed,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"actions": {
						Type:        schema.TypeSet,
						Description: "Actions is a list of actions this role binding grants access to.",
						Required:    !t.Computed,
						Computed:    t.Computed,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"paths": {
						Type:        schema.TypeSet,
						Description: "Paths is a list of paths this role binding grants access to.",
						Optional:    !t.Computed,
						Computed:    t.Computed,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"resources": {
						Type:        schema.TypeList,
						Description: "Resources is a list of resources this role binding grants access to.",
						Optional:    !t.Computed,
						Computed:    t.Computed,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"api_group": {
									Type:        schema.TypeString,
									Description: "APIGroup is the group for the resource being referenced.",
									Required:    !t.Computed,
									Computed:    t.Computed,
								},
								"kind": {
									Type:        schema.TypeString,
									Description: "Kind is the type of resource being referenced.",
									Required:    !t.Computed,
									Computed:    t.Computed,
								},
							},
						},
					},
				},
			},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.RoleSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)

	rawStatementList := fwflex.ExpandMapList(attr["statements"].([]any))
	for _, rawStatement := range rawStatementList {
		statement := bluechip_models.PolicyStatement{
			Actions: fwflex.ExpandStringSet(rawStatement["actions"].(*schema.Set)),
		}

		if rawStatement["paths"] != nil {
			statement.Paths = fwflex.ExpandStringSet(rawStatement["paths"].(*schema.Set))
		}

		if rawStatement["resources"] != nil {
			for _, rawPolicyResource := range fwflex.ExpandMapList(rawStatement["resources"].([]any)) {
				policyResource := bluechip_models.PolicyResource{
					ApiGroup: rawPolicyResource["api_group"].(string),
					Kind:     rawPolicyResource["kind"].(string),
				}
				statement.Resources = append(statement.Resources, policyResource)
			}
		}

		out.Statements = append(out.Statements, statement)
	}

	return nil
}

func (t SpecType) Flatten(in bluechip_models.RoleSpec) map[string]any {
	attr := map[string]any{
		"statements": []map[string]any{},
	}

	for _, statement := range in.Statements {
		rawStatement := map[string]any{
			"actions": statement.Actions,
		}

		if len(statement.Paths) != 0 {
			rawStatement["paths"] = statement.Paths
		}

		var policyResources []map[string]any
		for _, policyResource := range statement.Resources {
			rawPolicyResource := map[string]any{
				"api_group": policyResource.ApiGroup,
				"kind":      policyResource.Kind,
			}
			policyResources = append(policyResources, rawPolicyResource)
		}
		if len(policyResources) != 0 {
			rawStatement["resources"] = policyResources
		}
		attr["statements"] = append(attr["statements"].([]map[string]any), rawStatement)
	}
	return attr
}
