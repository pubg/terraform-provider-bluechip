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
	Computed bool
}

func (t SpecType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"ipv4_cidrs": {
			Type:     schema.TypeSet,
			Optional: !t.Computed,
			Computed: t.Computed,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
		"ipv6_cidrs": {
			Type:     schema.TypeSet,
			Optional: !t.Computed,
			Computed: t.Computed,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.IsCIDR,
			},
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.CidrSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.Ipv4Cidrs = fwflex.ExpandStringSet(attr["ipv4_cidrs"].(*schema.Set))
	out.Ipv6Cidrs = fwflex.ExpandStringSet(attr["ipv6_cidrs"].(*schema.Set))
	return nil
}

func (t SpecType) Flatten(in bluechip_models.CidrSpec) map[string]any {
	attr := map[string]any{
		"ipv4_cidrs": in.Ipv4Cidrs,
		"ipv6_cidrs": in.Ipv6Cidrs,
	}
	return attr
}
