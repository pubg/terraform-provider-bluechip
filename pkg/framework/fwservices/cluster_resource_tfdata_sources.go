package fwservices

import (
	"context"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type ClusterTerraformDataSources[T bluechip_models.NamespacedApiResource[P], P bluechip_models.BaseSpec] struct {
	Timeout time.Duration
	Gvk     bluechip_models.GroupVersionKind

	FilterType   fwtype.FilterType
	MetadataType fwtype.TypeHelper[bluechip_models.Metadata]
	SpecType     fwtype.TypeHelper[P]
}

func (r *ClusterTerraformDataSources[T, P]) Resource() *schema.Resource {
	resourceSchema := map[string]*schema.Schema{
		"metadata": r.MetadataType.Schema(),
	}
	if specSchema := r.SpecType.Schema(); specSchema != nil {
		resourceSchema["spec"] = specSchema
	}

	outerSchema := map[string]*schema.Schema{
		"filter": r.FilterType.Schema(),
		"items": {
			Type:        schema.TypeList,
			Description: "Filter is a list of query terms to filter the results by.",
			Computed:    true,
			Elem: &schema.Resource{
				Schema:      resourceSchema,
				Description: "Filter is a list of query terms to filter the results by.",
			},
		},
	}

	return &schema.Resource{
		Schema: outerSchema,
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Read(ctx, data, r.clusterWideClient(meta))
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(r.Timeout),
		},
	}
}

func (r *ClusterTerraformDataSources[T, P]) Read(ctx context.Context, d *schema.ResourceData, client *bluechip_client.ClusterResourceClient[T, P]) diag.Diagnostics {
	var filter fwtype.FilterValue
	if diags := r.FilterType.Expand(ctx, d, &filter); diags.HasError() {
		return diags
	}

	objects, err := r.listRequest(ctx, client, filter.Filter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("pubg")

	if diags := r.FilterType.FlattenWithSet(ctx, d, filter); diags.HasError() {
		return diags
	}

	var itemsAttr []map[string]any
	for _, object := range objects {
		metadataAttr := r.MetadataType.Flatten(object.GetMetadata())
		specAttr := r.SpecType.Flatten(object.GetSpec())

		itemAttr := map[string]any{"metadata": []any{metadataAttr}}
		if specAttr != nil {
			itemAttr["spec"] = []any{specAttr}
		}

		itemsAttr = append(itemsAttr, itemAttr)
	}

	if err := d.Set("items", itemsAttr); err != nil {
		return append(diag.FromErr(err), diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Failed to set items",
			Detail:        err.Error(),
			AttributePath: cty.GetAttrPath("items"),
		})
	}

	return nil
}

func (r *ClusterTerraformDataSources[T, P]) listRequest(ctx context.Context, client *bluechip_client.ClusterResourceClient[T, P], queryTerms []bluechip_models.QueryTerm) ([]T, error) {
	if len(queryTerms) == 0 {
		return client.List(ctx)
	} else {
		return client.Search(ctx, queryTerms)
	}
}

func (r *ClusterTerraformDataSources[T, P]) clusterWideClient(meta interface{}) *bluechip_client.ClusterResourceClient[T, P] {
	model := meta.(*provider.ProviderModel)
	return bluechip_client.NewClusterClient[T, P](model.Client, r.Gvk)
}
