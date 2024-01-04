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
}

func (SpecType) Schema(computed bool) *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"project": {
			Required: !computed,
			Computed: computed,
			Type:     schema.TypeString,
		},
		"environment": {
			Required: !computed,
			Computed: computed,
			Type:     schema.TypeString,
		},
		"organization_unit": {
			Required: !computed,
			Computed: computed,
			Type:     schema.TypeString,
		},
		"platform": {
			Required: !computed,
			Computed: computed,
			Type:     schema.TypeString,
		},
		"pubg": fwtype.SingleNestedBlock(map[string]*schema.Schema{
			"infra": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
			"site": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
		}, computed, true),
		"vendor": fwtype.SingleNestedBlock(map[string]*schema.Schema{
			"name": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
			"account_id": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
			"engine": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
			"region": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
		}, computed, false),
		"kubernetes": fwtype.SingleNestedBlock(map[string]*schema.Schema{
			"endpoint": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
			"ca_cert": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
			"sa_issuer": {
				Required:     !computed,
				Computed:     computed,
				Type:         schema.TypeString,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"version": {
				Required: !computed,
				Computed: computed,
				Type:     schema.TypeString,
			},
		}, computed, false),
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.ClusterSpec) diag.Diagnostics {
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

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.ClusterSpec) diag.Diagnostics {
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

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
