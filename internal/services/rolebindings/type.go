package rolebindings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type SpecType struct {
	RbType fwtype.RoleBindingType
}

func (t SpecType) Schema() *schema.Schema {
	return t.RbType.Schema()
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.RoleBindingSpec) diag.Diagnostics {
	return t.RbType.Expand(ctx, d, out)
}

func (t SpecType) Flatten(in bluechip_models.RoleBindingSpec) map[string]any {
	return t.RbType.Flatten(in)
}
