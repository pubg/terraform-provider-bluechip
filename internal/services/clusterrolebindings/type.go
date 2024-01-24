package clusterrolebindings

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

func (t SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.ClusterRoleBindingSpec) diag.Diagnostics {
	var rb bluechip_models.RoleBindingSpec
	if diags := t.RbType.Expand(ctx, d, &rb); diags.HasError() {
		return diags
	}

	out.SubjectsRef = rb.SubjectsRef
	out.RoleRef = rb.RoleRef
	return nil
}

func (t SpecType) Flatten(in bluechip_models.ClusterRoleBindingSpec) map[string]any {
	var rb bluechip_models.RoleBindingSpec
	rb.SubjectsRef = in.SubjectsRef
	rb.RoleRef = in.RoleRef

	return t.RbType.Flatten(rb)
}
