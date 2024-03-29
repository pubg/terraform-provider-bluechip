package fwtype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwflex"
)

var _ TypeHelper[bluechip_models.Metadata] = &MetadataType{}

type MetadataType struct {
	Namespaced bool
	Computed   bool
}

func NewMetadataType(namespaced bool, computed bool) *MetadataType {
	return &MetadataType{Namespaced: namespaced, Computed: computed}
}

func (t MetadataType) Schema() *schema.Schema {
	innerSchema := map[string]*schema.Schema{
		"annotations": {
			Type:        schema.TypeMap,
			Elem:        schema.TypeString,
			Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.",
			Computed:    t.Computed,
			Optional:    true,
		},
		"labels": {
			Type:        schema.TypeMap,
			Elem:        schema.TypeString,
			Description: "Labels are key value pairs that may be used to scope and select individual resources. They are not queryable and should be preserved when modifying objects.",
			Computed:    t.Computed,
			Optional:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Name is the name of the resource.",
			Required:    true,
		},
		"creation_timestamp": {
			Type:        schema.TypeString,
			Description: "CreationTimestamp is a timestamp representing the server time when this object was created.",
			Computed:    true,
		},
		"update_timestamp": {
			Type:        schema.TypeString,
			Description: "UpdateTimestamp is a timestamp representing the server time when this object was last updated.",
			Computed:    true,
		},
	}

	if t.Namespaced {
		innerSchema["namespace"] = &schema.Schema{
			Type:        schema.TypeString,
			Description: "Namespace is the namespace of the resource.",
			Required:    true,
		}
	}

	blockSchema := SingleNestedBlock(innerSchema, t.Computed, false)
	blockSchema.Optional = t.Computed
	CleanForDataSource(blockSchema)
	return blockSchema
}

func (MetadataType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.Metadata) diag.Diagnostics {
	attr, diags := fwflex.ExtractSingleBlock(ctx, d, "metadata")
	if diags.HasError() {
		return diags
	}

	out.Name = attr["name"].(string)
	if attr["namespace"] != nil {
		out.Namespace = attr["namespace"].(string)
	}
	if attr["annotations"] != nil {
		out.Annotations = fwflex.ExpandMap(attr["annotations"].(map[string]any))
	}
	if attr["labels"] != nil {
		out.Labels = fwflex.ExpandMap(attr["labels"].(map[string]any))
	}
	if attr["creation_timestamp"] != nil {
		out.CreationTimestamp = attr["creation_timestamp"].(string)
	}
	if attr["update_timestamp"] != nil {
		out.UpdateTimestamp = attr["update_timestamp"].(string)
	}
	return nil
}

func (MetadataType) Flatten(in bluechip_models.Metadata) map[string]any {
	attr := map[string]any{}
	attr["name"] = in.Name
	if in.Namespace != "" {
		attr["namespace"] = in.Namespace
	}
	attr["annotations"] = fwflex.FlattenMap(in.Annotations)
	attr["labels"] = fwflex.FlattenMap(in.Labels)
	if in.CreationTimestamp != "" {
		attr["creation_timestamp"] = in.CreationTimestamp
	}
	if in.UpdateTimestamp != "" {
		attr["update_timestamp"] = in.UpdateTimestamp
	}
	return attr
}
