package fwservices

import (
	"context"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwlog"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type ClusterTerraformResource[T bluechip_models.ClusterApiResource[P], P bluechip_models.BaseSpec] struct {
	Timeout     time.Duration
	Gvk         bluechip_models.GroupVersionKind
	Constructor func() T

	MetadataType fwtype.TypeHelper[bluechip_models.Metadata]
	SpecType     fwtype.TypeHelper[P]
}

func (r *ClusterTerraformResource[T, P]) Resource() *schema.Resource {
	scheme := map[string]*schema.Schema{
		"metadata": r.MetadataType.Schema(),
	}

	if specSchema := r.SpecType.Schema(); specSchema != nil {
		scheme["spec"] = specSchema
	}

	return &schema.Resource{
		Schema: scheme,
		CreateContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Upsert(ctx, data, r.clusterWideClient(meta))
		},
		ReadContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Read(ctx, data, r.clusterWideClient(meta))
		},
		UpdateContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Upsert(ctx, data, r.clusterWideClient(meta))
		},
		DeleteContext: func(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
			return r.Delete(ctx, data, r.clusterWideClient(meta))
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(r.Timeout),
		},
	}
}

func (r *ClusterTerraformResource[T, P]) Upsert(ctx context.Context, d *schema.ResourceData, client *bluechip_client.ClusterResourceClient[T, P]) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "gvk", r.Gvk)

	var metadata bluechip_models.Metadata
	if diags := r.MetadataType.Expand(ctx, d, &metadata); diags.HasError() {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "metadata Expansion Failed",
			Detail:        "metadata Expansion Failed",
			AttributePath: cty.GetAttrPath("metadata.0"),
		})
		tflog.Info(ctx, "metadata Expansion Failed", fwlog.Field("diags", diags))
		return diags
	}

	var spec P
	if diags := r.SpecType.Expand(ctx, d, &spec); diags.HasError() {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "spec Expansion Failed",
			Detail:        "spec Expansion Failed",
			AttributePath: cty.GetAttrPath("spec.0"),
		})
		tflog.Info(ctx, "spec Expansion Failed", fwlog.Field("diags", diags))
		return diags
	}

	object := r.Constructor()
	typeMeta := r.Gvk.ToTypeMeta()
	object.SetApiVersion(typeMeta.GetApiVersion())
	object.SetKind(typeMeta.GetKind())
	object.SetMetadata(metadata)
	object.SetSpec(spec)
	if err := client.Upsert(ctx, object); err != nil {
		diags := append(diag.FromErr(err), diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Upsert Request Failed",
			Detail:   "Upsert Request Failed",
		})
		tflog.Info(ctx, "Upsert Request Failed", fwlog.Field("object", object), fwlog.Field("diags", diags))
		return diags
	}

	d.SetId(ClusterResourceIdentity(metadata.Name))
	return r.Read(ctx, d, client)
}

func (r *ClusterTerraformResource[T, P]) Read(ctx context.Context, d *schema.ResourceData, client *bluechip_client.ClusterResourceClient[T, P]) diag.Diagnostics {
	name := ClusterResourceIdentityFrom(d.Id())
	object, err := client.Get(ctx, name)
	if err != nil {
		diags := append(diag.FromErr(err), diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Get Request Failed",
			Detail:   "Get Request Failed",
		})
		tflog.Info(ctx, "Get Request Failed", fwlog.Field("diags", diags))
		return diags
	}

	if diags := fwtype.SetBlock(d, "metadata", r.MetadataType.Flatten(object.GetMetadata())); diags.HasError() {
		return diags
	}
	if diags := fwtype.SetBlock(d, "spec", r.SpecType.Flatten(object.GetSpec())); diags.HasError() {
		return diags
	}

	return nil
}

func (r *ClusterTerraformResource[T, P]) Delete(ctx context.Context, d *schema.ResourceData, client *bluechip_client.ClusterResourceClient[T, P]) diag.Diagnostics {
	name := ClusterResourceIdentityFrom(d.Id())

	if err := client.Delete(ctx, name); err != nil {
		diags := append(diag.FromErr(err), diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Delete Request Failed",
			Detail:   "Delete Request Failed",
		})
		tflog.Info(ctx, "Delete Request Failed", fwlog.Field("diags", diags))
		return diags
	}
	return nil
}

func (r *ClusterTerraformResource[T, P]) clusterWideClient(meta interface{}) *bluechip_client.ClusterResourceClient[T, P] {
	model := meta.(*provider.ProviderModel)
	return bluechip_client.NewClusterClient[T, P](model.Client, r.Gvk)
}
