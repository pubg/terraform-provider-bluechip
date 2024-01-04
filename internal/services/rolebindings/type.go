package rolebindings

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
	return fwtype.RoleBindingTyp.Schema(computed)
}

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.RoleBindingSpec) diag.Diagnostics {
	return fwtype.RoleBindingTyp.Expand(ctx, d, out)
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.RoleBindingSpec) diag.Diagnostics {
	return fwtype.RoleBindingTyp.Flatten(ctx, d, in)
}
