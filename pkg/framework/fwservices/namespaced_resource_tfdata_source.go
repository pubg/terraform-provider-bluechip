package fwservices

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
)

type NamespacedTerraformDataSource[T bluechip_models.NamespacedApiResource[P], P bluechip_models.BaseSpec] struct {
	Schema        map[string]*schema.Schema
	Timeout       time.Duration
	Gvk           bluechip_models.GroupVersionKind
	SpecExpander  fwflex.Expander[P]
	SpecFlattener fwflex.Flattener[P]
}

func (r *NamespacedTerraformDataSource[T, P]) Resource() *schema.Resource {
	return &schema.Resource{
		Schema: r.Schema,
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Read(ctx, data, r.namespacedClient(meta))
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(r.Timeout),
		},
	}
}

func (r *NamespacedTerraformDataSource[T, P]) Read(ctx context.Context, d *schema.ResourceData, client *bluechip_client.NamespacedResourceClient[T, P]) diag.Diagnostics {
	var metadata bluechip_models.Metadata
	diags := metadataTyp.Expand(ctx, d, &metadata)
	if diags.HasError() {
		return diags
	}

	object, err := client.Get(ctx, metadata.Namespace, metadata.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(NamespacedResourceIdentity(metadata.Namespace, metadata.Name))

	if diags := metadataTyp.Flatten(ctx, d, object.GetMetadata()); diags.HasError() {
		return diags
	}

	if diags := r.SpecFlattener(ctx, d, object.GetSpec()); diags.HasError() {
		return diags
	}

	return nil
}

func (r *NamespacedTerraformDataSource[T, P]) namespacedClient(meta interface{}) *bluechip_client.NamespacedResourceClient[T, P] {
	model := meta.(*provider.ProviderModel)
	return bluechip_client.NewNamespacedClient[T, P](model.Client, r.Gvk)
}
