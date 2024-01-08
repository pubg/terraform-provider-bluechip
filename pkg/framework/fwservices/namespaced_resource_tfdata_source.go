package fwservices

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type NamespacedTerraformDataSource[T bluechip_models.NamespacedApiResource[P], P bluechip_models.BaseSpec] struct {
	Timeout time.Duration
	Gvk     bluechip_models.GroupVersionKind

	MetadataType fwtype.TypeHelper[bluechip_models.Metadata]
	SpecType     fwtype.TypeHelper[P]
}

func (r *NamespacedTerraformDataSource[T, P]) Resource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metadata": r.MetadataType.Schema(),
			"spec":     r.SpecType.Schema(),
		},
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
	if diags := r.MetadataType.Expand(ctx, d, &metadata); diags.HasError() {
		return diags
	}

	object, err := client.Get(ctx, metadata.Namespace, metadata.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(NamespacedResourceIdentity(metadata.Namespace, metadata.Name))

	if diags := fwtype.SetBlock(d, "metadata", r.MetadataType.Flatten(object.GetMetadata())); diags.HasError() {
		return diags
	}
	if diags := fwtype.SetBlock(d, "spec", r.SpecType.Flatten(object.GetSpec())); diags.HasError() {
		return diags
	}

	return nil
}

func (r *NamespacedTerraformDataSource[T, P]) namespacedClient(meta interface{}) *bluechip_client.NamespacedResourceClient[T, P] {
	model := meta.(*provider.ProviderModel)
	return bluechip_client.NewNamespacedClient[T, P](model.Client, r.Gvk)
}
