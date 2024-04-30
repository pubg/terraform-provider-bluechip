package fwservices

import (
	"context"
	"time"

	"git.projectbro.com/Devops/arcane-client-go/pkg/api_client"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"
	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwbuilder"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwlog"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ClusterTerraformResource[T api_meta.ClusterApiResource, S any] struct {
	Timeout time.Duration
	Gvk     api_meta.ExtendedGroupVersionKind

	MetadataType     fwtype.TypeHelper[api_meta.Metadata]
	SpecType         fwtype.TypeHelper[S]
	DebuilderFactory fwbuilder.ResourceDebuilderFactory[T]
	BuilderFactory   fwbuilder.ResourceBuilderFactory[T]
}

func (r *ClusterTerraformResource[T, S]) Resource() *schema.Resource {
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

func (r *ClusterTerraformResource[T, S]) Upsert(ctx context.Context, d *schema.ResourceData, client *api_client.ArcaneClusterResourceClient[T]) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "gvk", r.Gvk)

	var metadata api_meta.Metadata
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

	var spec S
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

	typeMeta := r.Gvk.ToTypeMeta()
	builder := r.BuilderFactory.New()
	if err := builder.Set(fwbuilder.FieldApiVersion, typeMeta.GetApiVersion()); err != nil {
		return diag.FromErr(err)
	}
	if err := builder.Set(fwbuilder.FieldKind, typeMeta.GetKind()); err != nil {
		return diag.FromErr(err)
	}
	if err := builder.Set(fwbuilder.FieldMetadata, metadata); err != nil {
		return diag.FromErr(err)
	}
	if err := builder.Set(fwbuilder.FieldSpec, spec); err != nil {
		return diag.FromErr(err)
	}

	object := builder.Build()
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

func (r *ClusterTerraformResource[T, S]) Read(ctx context.Context, d *schema.ResourceData, client *api_client.ArcaneClusterResourceClient[T]) diag.Diagnostics {
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

func (r *ClusterTerraformResource[T, S]) Delete(ctx context.Context, d *schema.ResourceData, client *api_client.ArcaneClusterResourceClient[T]) diag.Diagnostics {
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

func (r *ClusterTerraformResource[T, S]) clusterWideClient(meta interface{}) *api_client.ArcaneClusterResourceClient[T] {
	model := meta.(*provider.ProviderModel)
	return api_client.NewClusterResourceClient[T](model.Client.GetArcaneClient(), r.Gvk, model.Client.GetBasePath())
}
