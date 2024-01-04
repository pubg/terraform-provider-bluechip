package images

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type SpecType struct {
}

func (SpecType) Schema(computed bool) *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"app": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"timestamp": {
			Type:     schema.TypeInt,
			Required: !computed,
			Computed: computed,
		},
		"commit_hash": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"repository": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"tag": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
		"branch": {
			Type:     schema.TypeString,
			Required: !computed,
			Computed: computed,
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.ImageSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.App = attr["app"].(string)
	out.Timestamp = attr["timestamp"].(int)
	out.CommitHash = attr["commit_hash"].(string)
	out.Repository = attr["repository"].(string)
	out.Tag = attr["tag"].(string)
	out.Branch = attr["branch"].(string)
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.ImageSpec) diag.Diagnostics {
	attr := map[string]any{
		"app":         in.App,
		"timestamp":   in.Timestamp,
		"commit_hash": in.CommitHash,
		"repository":  in.Repository,
		"tag":         in.Tag,
		"branch":      in.Branch,
	}

	if err := d.Set("spec", []map[string]any{attr}); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
