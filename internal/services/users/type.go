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
}

func (SpecType) Schema(computed bool) *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"password": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"groups": {
			Type:     schema.TypeSet,
			Required: !computed,
			Computed: computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"attributes": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.UserSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.Password = attr["password"].(string)
	out.Groups = fwflex.ExpandStringSet(attr["groups"].(*schema.Set))
	out.Attributes = fwflex.ExpandMap(attr["attributes"].(map[string]any))
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.UserSpec) diag.Diagnostics {
	attr := map[string]any{
		"password":   in.Password,
		"groups":     in.Groups,
		"attributes": in.Attributes,
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
