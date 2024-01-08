package fwtype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TypeHelper[T any] interface {
	Schema() *schema.Schema
	Expand(ctx context.Context, d *schema.ResourceData, out *T) diag.Diagnostics
	Flatten(in T) map[string]any
}
