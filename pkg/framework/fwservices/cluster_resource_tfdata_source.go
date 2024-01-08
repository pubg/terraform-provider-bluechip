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

type ClusterTerraformDataSource[T bluechip_models.ClusterApiResource[P], P bluechip_models.BaseSpec] struct {
	Timeout time.Duration
	Gvk     bluechip_models.GroupVersionKind

	MetadataType fwtype.TypeHelper[bluechip_models.Metadata]
	SpecType     fwtype.TypeHelper[P]
}

func (r *ClusterTerraformDataSource[T, P]) Resource() *schema.Resource {
	scheme := map[string]*schema.Schema{
		"metadata": r.MetadataType.Schema(),
	}

	if specSchema := r.SpecType.Schema(); specSchema != nil {
		scheme["spec"] = specSchema
	}

	return &schema.Resource{
		Schema: scheme,
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Read(ctx, data, r.clusterWideClient(meta))
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(r.Timeout),
		},
	}
}

func (r *ClusterTerraformDataSource[T, P]) Read(ctx context.Context, d *schema.ResourceData, client *bluechip_client.ClusterResourceClient[T, P]) diag.Diagnostics {
	var metadata bluechip_models.Metadata
	if diags := r.MetadataType.Expand(ctx, d, &metadata); diags.HasError() {
		return diags
	}

	object, err := client.Get(ctx, metadata.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ClusterResourceIdentity(metadata.Name))

	if diags := fwtype.SetBlock(d, "metadata", r.MetadataType.Flatten(object.GetMetadata())); diags.HasError() {
		return diags
	}
	if diags := fwtype.SetBlock(d, "spec", r.SpecType.Flatten(object.GetSpec())); diags.HasError() {
		return diags
	}

	return nil
}

func (r *ClusterTerraformDataSource[T, P]) clusterWideClient(meta interface{}) *bluechip_client.ClusterResourceClient[T, P] {
	model := meta.(*provider.ProviderModel)
	return bluechip_client.NewClusterClient[T, P](model.Client, r.Gvk)
}
