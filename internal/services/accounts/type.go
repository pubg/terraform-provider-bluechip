package accounts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

var _ fwtype.TypeHelper[bluechip_models.AccountSpec] = &SpecType{}

type SpecType struct {
	Computed bool
}

func (t SpecType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"display_name": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"description": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"alias": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"vendor": {
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

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.AccountSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.AccountId = attr["account_id"].(string)
	out.DisplayName = attr["display_name"].(string)
	out.Description = attr["description"].(string)
	out.Alias = attr["alias"].(string)
	out.Vendor = attr["vendor"].(string)
	out.Regions = fwflex.ExpandStringSet(attr["regions"].(*schema.Set))
	return nil
}

func (t SpecType) Flatten(in bluechip_models.AccountSpec) map[string]any {
	attr := map[string]any{
		"account_id":   in.AccountId,
		"display_name": in.DisplayName,
		"description":  in.Description,
		"alias":        in.Alias,
		"vendor":       in.Vendor,
		"regions":      in.Regions,
	}
	return attr
}
