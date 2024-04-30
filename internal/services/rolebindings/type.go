package rolebindings

import (
	"context"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SpecType struct {
	RbType fwtype.RoleBindingType
}

func (t SpecType) Schema() *schema.Schema {
	return t.RbType.Schema()
}

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip.RoleBindingSpec) diag.Diagnostics {
	return t.RbType.Expand(ctx, d, out)
}

func (t SpecType) Flatten(in bluechip.RoleBindingSpec) map[string]any {
	return t.RbType.Flatten(in)
}
