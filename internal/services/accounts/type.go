package accounts

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
		"account_id": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"display_name": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"description": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"alias": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"vendor": {
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

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.AccountSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.AccountId = attr["account_id"].(string)
	out.DisplayName = attr["display_name"].(string)
	out.Description = attr["description"].(string)
	out.Alias = attr["alias"].(string)
	out.Vendor = attr["vendor"].(string)
	out.Regions = fwflex.ExpandStringSet(attr["regions"].(*schema.Set))
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.AccountSpec) diag.Diagnostics {
	attr := map[string]any{
		"account_id":   in.AccountId,
		"display_name": in.DisplayName,
		"description":  in.Description,
		"alias":        in.Alias,
		"vendor":       in.Vendor,
		"regions":      in.Regions,
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
