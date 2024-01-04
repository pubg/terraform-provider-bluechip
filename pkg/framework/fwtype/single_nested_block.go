package fwtype

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func SingleNestedBlock(in map[string]*schema.Schema, computed bool, optional bool) *schema.Schema {
	block := &schema.Schema{
		Type:     schema.TypeList,
		Required: !computed,
		Computed: computed,
		Elem:     &schema.Resource{Schema: in},
	}

	if !computed {
		block.MaxItems = 1
		if optional {
			block.MinItems = 0
		} else {
			block.MinItems = 1
		}
	}
	return block
}
