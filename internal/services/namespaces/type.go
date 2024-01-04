package namespaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
)

type SpecType struct {
}

func (SpecType) Schema(computed bool) *schema.Schema {
	return nil
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.EmptySpec) diag.Diagnostics {
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.EmptySpec) diag.Diagnostics {
	return nil
}
