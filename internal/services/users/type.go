package users

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
		"password": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"groups": {
			Type:     schema.TypeSet,
			Optional: !t.Computed,
			Computed: t.Computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"attributes": {
			Type:     schema.TypeMap,
			Required: !t.Computed,
			Computed: t.Computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip.UserSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.Password = attr["password"].(string)
	if attr["groups"] != nil {
		out.Groups = fwflex.ExpandStringSet(attr["groups"].(*schema.Set))
	}
	out.Attributes = fwflex.ExpandMap(attr["attributes"].(map[string]any))
	return nil
}

func (t SpecType) Flatten(in bluechip.UserSpec) map[string]any {
	attr := map[string]any{
		//"password":   in.Password,
		"groups":     in.Groups,
		"attributes": in.Attributes,
	}

	return attr
}
