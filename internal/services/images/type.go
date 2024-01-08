package images

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type SpecType struct {
	Computed bool
}

func (t SpecType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"app": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"timestamp": {
			Type:     schema.TypeInt,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"commit_hash": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"repository": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"tag": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
		"branch": {
			Type:     schema.TypeString,
			Required: !t.Computed,
			Computed: t.Computed,
		},
	}

	blockSchema := fwtype.SingleNestedBlock(innerSchema, t.Computed, true)
	fwtype.CleanForDataSource(blockSchema)
	return blockSchema
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.ImageSpec) diag.Diagnostics {
	attr := d.Get("spec.0").(map[string]any)
	out.App = attr["app"].(string)
	out.Timestamp = attr["timestamp"].(int)
	out.CommitHash = attr["commit_hash"].(string)
	out.Repository = attr["repository"].(string)
	out.Tag = attr["tag"].(string)
	out.Branch = attr["branch"].(string)
	return nil
}

func (t SpecType) Flatten(in bluechip_models.ImageSpec) map[string]any {
	attr := map[string]any{
		"app":         in.App,
		"timestamp":   in.Timestamp,
		"commit_hash": in.CommitHash,
		"repository":  in.Repository,
		"tag":         in.Tag,
		"branch":      in.Branch,
	}

	return attr
}
