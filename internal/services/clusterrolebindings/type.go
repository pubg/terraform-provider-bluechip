package clusterrolebindings

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

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip.ClusterRoleBindingSpec) diag.Diagnostics {
	var rb bluechip.RoleBindingSpec
	if diags := t.RbType.Expand(ctx, d, &rb); diags.HasError() {
		return diags
	}

	out.SubjectsRef = rb.SubjectsRef
	out.RoleRef = rb.RoleRef
	return nil
}

func (t SpecType) Flatten(in bluechip.ClusterRoleBindingSpec) map[string]any {
	var rb bluechip.RoleBindingSpec
	rb.SubjectsRef = in.SubjectsRef
	rb.RoleRef = in.RoleRef

	return t.RbType.Flatten(rb)
}
