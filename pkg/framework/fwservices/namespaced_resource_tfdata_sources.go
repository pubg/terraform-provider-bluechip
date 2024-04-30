package fwservices

import (
	"context"
	"time"

	"git.projectbro.com/Devops/arcane-client-go/pkg/api_client"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"
	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwbuilder"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NamespacedTerraformDataSources[T api_meta.NamespacedApiResource, S any] struct {
	Timeout time.Duration
	Gvk     api_meta.ExtendedGroupVersionKind

	FilterType       fwtype.FilterType
	MetadataType     fwtype.TypeHelper[api_meta.Metadata]
	SpecType         fwtype.TypeHelper[S]
	DebuilderFactory fwbuilder.ResourceDebuilderFactory[T]
}

func (r *NamespacedTerraformDataSources[T, S]) Resource() *schema.Resource {
	resourceSchema := map[string]*schema.Schema{
		"metadata": r.MetadataType.Schema(),
	}
	if specSchema := r.SpecType.Schema(); specSchema != nil {
		resourceSchema["spec"] = specSchema
	}

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
				Schema:      resourceSchema,
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

func (r *NamespacedTerraformDataSources[T, S]) Read(ctx context.Context, d *schema.ResourceData, client *api_client.ArcaneNamespacedResourceClient[T]) diag.Diagnostics {
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
		debuilder := r.DebuilderFactory.New(object)

		metadata := debuilder.Get(fwbuilder.FieldMetadata).(api_meta.Metadata)
		metadataAttr := r.MetadataType.Flatten(metadata)
		itemAttr := map[string]any{"metadata": []any{metadataAttr}}

		spec := debuilder.Get(fwbuilder.FieldSpec).(S)
		specAttr := r.SpecType.Flatten(spec)
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

func (r *NamespacedTerraformDataSources[T, S]) listRequest(ctx context.Context, client *api_client.ArcaneNamespacedResourceClient[T], namespace string, queryTerms []api_meta.QueryTerm) ([]T, error) {
	if len(queryTerms) == 0 {
		return client.List(ctx, namespace)
	} else {
		return client.Search(ctx, namespace, queryTerms)
	}
}

func (r *NamespacedTerraformDataSources[T, S]) namespacedClient(meta interface{}) *api_client.ArcaneNamespacedResourceClient[T] {
	model := meta.(*provider.ProviderModel)
	return api_client.NewNamespacedResourceClient[T](model.Client.GetArcaneClient(), r.Gvk, model.Client.GetBasePath())
}
