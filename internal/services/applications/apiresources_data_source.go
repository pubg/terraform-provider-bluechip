package applications

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
)

func NewApiResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApiResourcesRead,
		Schema: map[string]*schema.Schema{
			"api_resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind_plural": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespaced": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(5 * time.Second),
		},
	}
}

func dataSourceApiResourcesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	model := meta.(*provider.ProviderModel)
	apiResourcesResp, err := model.Client.ApiResources(ctx)
	if err != nil {
		return diag.Errorf("Unable to read example, got error: %s", err)
	}

	d.SetId("api-resources")

	var apiResources []map[string]any
	for _, apiResource := range apiResourcesResp.Items {
		apiResources = append(apiResources, map[string]any{
			"group":       apiResource.Group,
			"version":     apiResource.Version,
			"kind":        apiResource.Kind,
			"kind_plural": apiResource.KindPlural,
			"namespaced":  apiResource.Namespaced,
		})
	}
	if err = d.Set("api_resources", apiResources); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
