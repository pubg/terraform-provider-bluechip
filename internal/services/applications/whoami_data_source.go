package applications

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
)

func NewWhoamiDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWhoamiRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"attributes": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(5 * time.Second),
		},
	}
}

func dataSourceWhoamiRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := meta.(*provider.ProviderModel)
	whoamiResp, err := model.Client.Whoami(ctx)
	if err != nil {
		return diag.Errorf("Unable to read example, got error: %s", err)
	}

	d.SetId(whoamiResp.Name)
	if err := d.Set("name", whoamiResp.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("groups", whoamiResp.Groups); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("attributes", whoamiResp.Attributes); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
