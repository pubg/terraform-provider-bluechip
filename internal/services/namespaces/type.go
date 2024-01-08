package namespaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
)

type SpecType struct {
}

func (t SpecType) Schema() *schema.Schema {
	return nil
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.EmptySpec) diag.Diagnostics {
	return nil
}

func (t SpecType) Flatten(in bluechip_models.EmptySpec) map[string]any {
	return nil
}
