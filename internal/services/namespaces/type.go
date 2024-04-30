package namespaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SpecType struct {
}

func (t SpecType) Schema() *schema.Schema {
	return nil
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *EmptySpec) diag.Diagnostics {
	return nil
}

func (t SpecType) Flatten(in EmptySpec) map[string]any {
	return nil
}
