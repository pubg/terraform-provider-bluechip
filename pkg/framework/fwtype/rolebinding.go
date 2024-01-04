package fwtype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
)

var RoleBindingTyp = &RoleBindingType{}

type RoleBindingType struct {
}

func (RoleBindingType) Schema(computed bool) *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"subject_ref": SingleNestedBlock(map[string]*schema.Schema{
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind of the referent. Valid kinds are 'User', 'Group'.",
				Required:    !computed,
				Computed:    computed,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the referent.",
				Required:    !computed,
				Computed:    computed,
			},
		}, computed, true),
		"policy_inline": {
			Type:     schema.TypeList,
			Optional: !computed,
			Computed: computed,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"actions": {
						Type:        schema.TypeSet,
						Description: "Actions is a list of actions this role binding grants access to.",
						Required:    !computed,
						Computed:    computed,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"paths": {
						Type:        schema.TypeSet,
						Description: "Paths is a list of paths this role binding grants access to.",
						Required:    !computed,
						Computed:    computed,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					"resources": {
						Type:        schema.TypeSet,
						Description: "Resources is a list of resources this role binding grants access to.",
						Required:    !computed,
						Computed:    computed,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"api_group": {
									Type:        schema.TypeString,
									Description: "APIGroup is the group for the resource being referenced.",
									Required:    !computed,
									Computed:    computed,
								},
								"kind": {
									Type:        schema.TypeString,
									Description: "Kind is the type of resource being referenced.",
									Required:    !computed,
									Computed:    computed,
								},
							},
						},
					},
				},
			},
		},
		"policy_ref": {
			Type:     schema.TypeString,
			Optional: !computed,
			Computed: computed,
		},
	}

	blockSchema := SingleNestedBlock(innerSchema, computed, true)
	CleanForDataSource(blockSchema)
	return blockSchema
}

func (RoleBindingType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.RoleBindingSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)

	rawSubjectRef := fwflex.ExpandMapList(attr["subject_ref"].([]any))
	out.SubjectsRef = bluechip_models.SubjectRef{
		Kind: rawSubjectRef[0]["kind"].(string),
		Name: rawSubjectRef[0]["name"].(string),
	}

	rawPolicyInlineList := fwflex.ExpandMapList(attr["policy_inline"].([]any))
	for _, rawPolicyInline := range rawPolicyInlineList {
		policyInline := bluechip_models.PolicyStatement{
			Actions:   fwflex.ExpandStringSet(rawPolicyInline["actions"].(*schema.Set)),
			Paths:     fwflex.ExpandStringSet(rawPolicyInline["paths"].(*schema.Set)),
			Resources: []bluechip_models.PolicyResource{},
		}
		rawPolicyResourceList := fwflex.ExpandMapList(rawPolicyInline["resources"].([]any))
		for _, rawPolicyResource := range rawPolicyResourceList {
			policyResource := bluechip_models.PolicyResource{
				ApiGroup: rawPolicyResource["api_group"].(string),
				Kind:     rawPolicyResource["kind"].(string),
			}
			policyInline.Resources = append(policyInline.Resources, policyResource)
		}
		out.PolicyInline = append(out.PolicyInline, policyInline)
	}

	if attr["policy_ref"] != nil {
		out.PolicyRef = String(attr["policy_ref"].(string))
	}
	return nil
}

func (RoleBindingType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.RoleBindingSpec) diag.Diagnostics {
	attr := map[string]any{
		"subject_ref": []map[string]any{{
			"kind": in.SubjectsRef.Kind,
			"name": in.SubjectsRef.Name,
		}},
		"policy_inline": []map[string]any{},
		"policy_ref":    in.PolicyRef,
	}

	for _, policyInline := range in.PolicyInline {
		rawPolicyInline := map[string]any{
			"actions":   policyInline.Actions,
			"paths":     policyInline.Paths,
			"resources": []map[string]any{},
		}
		for _, policyResource := range policyInline.Resources {
			rawPolicyResource := map[string]any{
				"api_group": policyResource.ApiGroup,
				"kind":      policyResource.Kind,
			}
			rawPolicyInline["resources"] = append(rawPolicyInline["resources"].([]map[string]any), rawPolicyResource)
		}
		attr["policy_inline"] = append(attr["policy_inline"].([]map[string]any), rawPolicyInline)
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
