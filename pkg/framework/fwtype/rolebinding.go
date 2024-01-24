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
		"role_ref": SingleNestedBlock(map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the referent.",
				Required:    !t.Computed,
				Computed:    t.Computed,
			},
		}, t.Computed, true),
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

	rawRoleRef := fwflex.ExpandMapList(attr["role_ref"].([]any))
	out.RoleRef = bluechip_models.RoleRef{
		Name: rawRoleRef[0]["name"].(string),
	}
	return nil
}

func (RoleBindingType) Flatten(in bluechip_models.RoleBindingSpec) map[string]any {
	attr := map[string]any{
		"subject_ref": []map[string]any{{
			"kind": in.SubjectsRef.Kind,
			"name": in.SubjectsRef.Name,
		}},
		"role_ref": []map[string]any{{
			"name": in.RoleRef.Name,
		}},
	}
	return attr
}
