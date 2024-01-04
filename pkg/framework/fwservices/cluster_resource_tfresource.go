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
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwlog"
)

type ClusterTerraformResource[T bluechip_models.ClusterApiResource[P], P bluechip_models.BaseSpec] struct {
	Schema        map[string]*schema.Schema
	Timeout       time.Duration
	Gvk           bluechip_models.GroupVersionKind
	SpecExpander  fwflex.Expander[P]
	SpecFlattener fwflex.Flattener[P]
	Constructor   func() T
}

func (r *ClusterTerraformResource[T, P]) Resource() *schema.Resource {
	return &schema.Resource{
		Schema: r.Schema,
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
	if diags := metadataTyp.Expand(ctx, d, &metadata); diags.HasError() {
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
	if diags := r.SpecExpander(ctx, d, &spec); diags.HasError() {
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

	if diags := metadataTyp.Flatten(ctx, d, object.GetMetadata()); diags.HasError() {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "metadata Flatten Failed",
			Detail:        "metadata Flatten Failed",
			AttributePath: cty.GetAttrPath("metadata.0"),
		})
		tflog.Info(ctx, "metadata Flatten Failed", fwlog.Field("diags", diags))
		return diags
	}

	if diags := r.SpecFlattener(ctx, d, object.GetSpec()); diags.HasError() {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "spec Flatten Failed",
			Detail:        "spec Flatten Failed",
			AttributePath: cty.GetAttrPath("spec.0"),
		})
		tflog.Info(ctx, "spec Flatten Failed", fwlog.Field("diags", diags))
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
