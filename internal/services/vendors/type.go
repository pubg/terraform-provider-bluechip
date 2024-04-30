package vendors

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
		"display_name": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"code_name": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"short_name": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"regions": {
			Type:     schema.TypeSet,
			Required: !t.Computed,
			Computed: t.Computed,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip.VendorSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.DisplayName = attr["display_name"].(string)
	out.CodeName = attr["code_name"].(string)
	out.ShortName = attr["short_name"].(string)
	out.Regions = fwflex.ExpandStringSet(attr["regions"].(*schema.Set))
	return nil
}

func (t SpecType) Flatten(in bluechip.VendorSpec) map[string]any {
	attr := map[string]any{
		"display_name": in.DisplayName,
		"code_name":    in.CodeName,
		"short_name":   in.ShortName,
		"regions":      in.Regions,
	}

	return attr
}
