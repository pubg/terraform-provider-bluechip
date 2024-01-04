package vendors

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
		"display_name": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"code_name": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"short_name": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"regions": {
			Type:     schema.TypeSet,
			Required: !computed,
			Computed: computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.VendorSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.DisplayName = attr["display_name"].(string)
	out.CodeName = attr["code_name"].(string)
	out.ShortName = attr["short_name"].(string)
	out.Regions = fwflex.ExpandStringSet(attr["regions"].(*schema.Set))
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.VendorSpec) diag.Diagnostics {
	attr := map[string]any{
		"display_name": in.DisplayName,
		"code_name":    in.CodeName,
		"short_name":   in.ShortName,
		"regions":      in.Regions,
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
