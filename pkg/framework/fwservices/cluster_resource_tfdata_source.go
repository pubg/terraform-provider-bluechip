package fwservices

import (
	"context"
	"time"

	"git.projectbro.com/Devops/arcane-client-go/pkg/api_client"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"
	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwbuilder"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ClusterTerraformDataSource[T api_meta.ClusterApiResource, S any] struct {
	Timeout time.Duration
	Gvk     api_meta.ExtendedGroupVersionKind

	MetadataType     fwtype.TypeHelper[api_meta.Metadata]
	SpecType         fwtype.TypeHelper[S]
	DebuilderFactory fwbuilder.ResourceDebuilderFactory[T]
}

func (r *ClusterTerraformDataSource[T, S]) Resource() *schema.Resource {
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

func (r *ClusterTerraformDataSource[T, S]) Read(ctx context.Context, d *schema.ResourceData, client *api_client.ArcaneClusterResourceClient[T]) diag.Diagnostics {
	var requestMetadata api_meta.Metadata
	if diags := r.MetadataType.Expand(ctx, d, &requestMetadata); diags.HasError() {
		return diags
	}

	object, err := client.Get(ctx, requestMetadata.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ClusterResourceIdentity(requestMetadata.Name))

	debuilder := r.DebuilderFactory.New(object)

	metadata := debuilder.Get(fwbuilder.FieldMetadata).(api_meta.Metadata)
	if diags := fwtype.SetBlock(d, "metadata", r.MetadataType.Flatten(metadata)); diags.HasError() {
		return diags
	}
	if debuilder.Get(fwbuilder.FieldSpec) != nil {
		spec := debuilder.Get(fwbuilder.FieldSpec).(S)
		if diags := fwtype.SetBlock(d, "spec", r.SpecType.Flatten(spec)); diags.HasError() {
			return diags
		}
	}

	return nil
}

func (r *ClusterTerraformDataSource[T, S]) clusterWideClient(meta interface{}) *api_client.ArcaneClusterResourceClient[T] {
	model := meta.(*provider.ProviderModel)
	return api_client.NewClusterResourceClient[T](model.Client.GetArcaneClient(), r.Gvk, model.Client.GetBasePath())
}
