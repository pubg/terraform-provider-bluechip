package fwtype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
)

var _ TypeHelper[bluechip_models.RoleBindingSpec] = &RoleBindingType{}

type RoleBindingType struct {
	Computed bool
}

func (t RoleBindingType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"subject_ref": SingleNestedBlock(map[string]*schema.Schema{
			"kind": {
				Type:        schema.TypeString,
				Description: "Kind of the referent. Valid kinds are 'User', 'Group'.",
				Required:    !t.Computed,
				Computed:    t.Computed,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the referent.",
				Required:    !t.Computed,
				Computed:    t.Computed,
			},
		}, t.Computed, true),
		"policy_inline": {
			Type:     schema.TypeList,
			Optional: !t.Computed,
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
		"policy_ref": {
			Type:     schema.TypeString,
			Optional: !t.Computed,
			Computed: t.Computed,
		},
	}

	blockSchema := SingleNestedBlock(innerSchema, t.Computed, true)
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
			Actions: fwflex.ExpandStringSet(rawPolicyInline["actions"].(*schema.Set)),
		}

		if rawPolicyInline["paths"] != nil {
			policyInline.Paths = fwflex.ExpandStringSet(rawPolicyInline["paths"].(*schema.Set))
		}

		if rawPolicyInline["resources"] != nil {
			for _, rawPolicyResource := range fwflex.ExpandMapList(rawPolicyInline["resources"].([]any)) {
				policyResource := bluechip_models.PolicyResource{
					ApiGroup: rawPolicyResource["api_group"].(string),
					Kind:     rawPolicyResource["kind"].(string),
				}
				policyInline.Resources = append(policyInline.Resources, policyResource)
			}
		}

		out.PolicyInline = append(out.PolicyInline, policyInline)
	}

	if attr["policy_ref"] != nil {
		if attr["policy_ref"].(string) != "" {
			out.PolicyRef = String(attr["policy_ref"].(string))
		}
	}
	return nil
}

func (RoleBindingType) Flatten(in bluechip_models.RoleBindingSpec) map[string]any {
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
			"actions": policyInline.Actions,
		}

		if len(policyInline.Paths) != 0 {
			rawPolicyInline["paths"] = policyInline.Paths
		}

		var policyResources []map[string]any
		for _, policyResource := range policyInline.Resources {
			rawPolicyResource := map[string]any{
				"api_group": policyResource.ApiGroup,
				"kind":      policyResource.Kind,
			}
			policyResources = append(policyResources, rawPolicyResource)
		}
		if len(policyResources) != 0 {
			rawPolicyInline["resources"] = policyResources
		}
		attr["policy_inline"] = append(attr["policy_inline"].([]map[string]any), rawPolicyInline)
	}
	return attr
}
