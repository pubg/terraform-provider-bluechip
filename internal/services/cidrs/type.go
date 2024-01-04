package cidrs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type SpecType struct {
}

func (SpecType) Schema(computed bool) *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"ipv4_cidrs": {
			Type:     schema.TypeSet,
			Optional: !computed,
			Computed: computed,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		"ipv6_cidrs": {
			Type:     schema.TypeSet,
			Optional: !computed,
			Computed: computed,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.CidrSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.Ipv4Cidrs = fwflex.ExpandStringSet(attr["ipv4_cidrs"].(*schema.Set))
	out.Ipv6Cidrs = fwflex.ExpandStringSet(attr["ipv6_cidrs"].(*schema.Set))
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.CidrSpec) diag.Diagnostics {
	attr := map[string]any{
		"ipv4_cidrs": in.Ipv4Cidrs,
		"ipv6_cidrs": in.Ipv6Cidrs,
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
