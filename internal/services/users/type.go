package users

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
			Optional: true,
			Computed: t.Computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.UserSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.Password = attr["password"].(string)
	if attr["groups"] != nil {
		out.Groups = fwflex.ExpandStringSet(attr["groups"].(*schema.Set))
	}
	out.Attributes = fwflex.ExpandMap(attr["attributes"].(map[string]any))
	return nil
}

func (t SpecType) Flatten(in bluechip_models.UserSpec) map[string]any {
	attr := map[string]any{
		//"password":   in.Password,
		"groups":     in.Groups,
		"attributes": in.Attributes,
	}

	return attr
}
