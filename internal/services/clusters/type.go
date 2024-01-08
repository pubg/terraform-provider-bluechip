package clusters

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type SpecType struct {
	Computed bool
}

func (t SpecType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"project": {
			Required: !t.Computed,
			Computed: t.Computed,
			Type:     schema.TypeString,
		},
		"environment": {
			Required: !t.Computed,
			Computed: t.Computed,
			Type:     schema.TypeString,
		},
		"organization_unit": {
			Required: !t.Computed,
			Computed: t.Computed,
			Type:     schema.TypeString,
		},
		"platform": {
			Required: !t.Computed,
			Computed: t.Computed,
			Type:     schema.TypeString,
		},
		"pubg": fwtype.SingleNestedBlock(map[string]*schema.Schema{
			"infra": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
			"site": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
		}, t.Computed, true),
		"vendor": fwtype.SingleNestedBlock(map[string]*schema.Schema{
			"name": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
			"account_id": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
			"engine": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
			"region": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
		}, t.Computed, false),
		"kubernetes": fwtype.SingleNestedBlock(map[string]*schema.Schema{
			"endpoint": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
			"ca_cert": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
			"sa_issuer": {
				Required:     !t.Computed,
				Computed:     t.Computed,
				Type:         schema.TypeString,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"version": {
				Required: !t.Computed,
				Computed: t.Computed,
				Type:     schema.TypeString,
			},
		}, t.Computed, false),
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.ClusterSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.Project = attr["project"].(string)
	out.Environment = attr["environment"].(string)
	out.OrganizationUnit = attr["organization_unit"].(string)
	out.Platform = attr["platform"].(string)

	if pubgAttr, ok := d.GetOk("spec.0.pubg.0"); ok {
		pubg := fwflex.ExpandMap(pubgAttr.(map[string]any))
		out.Pubg = &bluechip_models.ClusterSpecPubg{
			Infra: pubg["infra"],
			Site:  pubg["site"],
		}
	}

	vendor := fwflex.ExpandMap(d.Get("spec.0.vendor.0").(map[string]any))
	out.Vendor = bluechip_models.ClusterSpecVendor{
		Name:      vendor["name"],
		AccountId: vendor["account_id"],
		Engine:    vendor["engine"],
		Region:    vendor["region"],
	}

	kubernetes := fwflex.ExpandMap(d.Get("spec.0.kubernetes.0").(map[string]any))
	out.Kubernetes = bluechip_models.ClusterSpecKubernetes{
		Endpoint: kubernetes["endpoint"],
		CaCert:   kubernetes["ca_cert"],
		SaIssuer: kubernetes["sa_issuer"],
		Version:  kubernetes["version"],
	}
	return nil
}

func (t SpecType) Flatten(in bluechip_models.ClusterSpec) map[string]any {
	attr := map[string]any{}
	attr["project"] = in.Project
	attr["environment"] = in.Environment
	attr["organization_unit"] = in.OrganizationUnit
	attr["platform"] = in.Platform

	if in.Pubg != nil {
		attr["pubg"] = []map[string]any{fwflex.FlattenMap(map[string]string{
			"infra": in.Pubg.Infra,
			"site":  in.Pubg.Site,
		})}
	}

	attr["vendor"] = []map[string]any{fwflex.FlattenMap(map[string]string{
		"name":       in.Vendor.Name,
		"account_id": in.Vendor.AccountId,
		"engine":     in.Vendor.Engine,
		"region":     in.Vendor.Region,
	})}

	attr["kubernetes"] = []map[string]any{fwflex.FlattenMap(map[string]string{
		"endpoint":  in.Kubernetes.Endpoint,
		"ca_cert":   in.Kubernetes.CaCert,
		"sa_issuer": in.Kubernetes.SaIssuer,
		"version":   in.Kubernetes.Version,
	})}

	return attr
}
