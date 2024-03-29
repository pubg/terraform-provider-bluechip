package fwflex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
)

type Flattener[T bluechip_models.BaseSpec] func(ctx context.Context, d *schema.ResourceData, in T) diag.Diagnostics

func FlattenMap(m map[string]string) map[string]any {
	attr := make(map[string]any)
	for k, v := range m {
		attr[k] = v
	}
	return attr
}
