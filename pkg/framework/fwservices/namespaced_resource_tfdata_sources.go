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

type NamespacedTerraformDataSources[T bluechip_models.NamespacedApiResource[P], P bluechip_models.BaseSpec] struct {
	Timeout time.Duration
	Gvk     bluechip_models.GroupVersionKind

	FilterType   fwtype.FilterType
	MetadataType fwtype.TypeHelper[bluechip_models.Metadata]
	SpecType     fwtype.TypeHelper[P]
}

func (r *NamespacedTerraformDataSources[T, P]) Resource() *schema.Resource {
	outerSchema := map[string]*schema.Schema{
		"filter": r.FilterType.Schema(),
		"namespace": {
			Type:        schema.TypeString,
			Description: "Namespace is the namespace of the resource.",
			Required:    true,
		},
		"items": {
			Type:        schema.TypeList,
			Description: "Filter is a list of query terms to filter the results by.",
			Computed:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"metadata": r.MetadataType.Schema(),
					"spec":     r.SpecType.Schema(),
				},
				Description: "Filter is a list of query terms to filter the results by.",
			},
		},
	}

	return &schema.Resource{
		Schema: outerSchema,
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Read(ctx, data, r.namespacedClient(meta))
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(r.Timeout),
		},
	}
}

func (r *NamespacedTerraformDataSources[T, P]) Read(ctx context.Context, d *schema.ResourceData, client *bluechip_client.NamespacedResourceClient[T, P]) diag.Diagnostics {
	var filter fwtype.FilterValue
	if diags := r.FilterType.Expand(ctx, d, &filter); diags.HasError() {
		return diags
	}

	namespace := d.Get("namespace").(string)
	objects, err := r.listRequest(ctx, client, namespace, filter.Filter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(namespace)

	if diags := r.FilterType.FlattenWithSet(ctx, d, filter); diags.HasError() {
		return diags
	}

	var itemsAttr []map[string]any
	for _, object := range objects {
		metadataAttr := r.MetadataType.Flatten(object.GetMetadata())
		specAttr := r.SpecType.Flatten(object.GetSpec())

		itemsAttr = append(itemsAttr, map[string]any{
			"metadata": []any{metadataAttr},
			"spec":     []any{specAttr},
		})
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

func (r *NamespacedTerraformDataSources[T, P]) listRequest(ctx context.Context, client *bluechip_client.NamespacedResourceClient[T, P], namespace string, queryTerms []bluechip_models.QueryTerm) ([]T, error) {
	if len(queryTerms) == 0 {
		return client.List(ctx, namespace)
	} else {
		return client.Search(ctx, namespace, queryTerms)
	}
}

func (r *NamespacedTerraformDataSources[T, P]) namespacedClient(meta interface{}) *bluechip_client.NamespacedResourceClient[T, P] {
	model := meta.(*provider.ProviderModel)
	return bluechip_client.NewNamespacedClient[T, P](model.Client, r.Gvk)
}
