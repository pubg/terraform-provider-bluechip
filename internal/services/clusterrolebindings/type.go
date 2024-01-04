package clusterrolebindings

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

func (SpecType) Expand(ctx context.Context, d *schema.ResourceData, out *bluechip_models.ClusterRoleBindingSpec) diag.Diagnostics {
	var rb bluechip_models.RoleBindingSpec
	diags := fwtype.RoleBindingTyp.Expand(ctx, d, &rb)
	if diags.HasError() {
		return diags
	}

	out.SubjectsRef = rb.SubjectsRef
	out.PolicyInline = rb.PolicyInline
	out.PolicyRef = rb.PolicyRef
	return nil
}

func (SpecType) Flatten(ctx context.Context, d *schema.ResourceData, in bluechip_models.ClusterRoleBindingSpec) diag.Diagnostics {
	var rb bluechip_models.RoleBindingSpec
	rb.SubjectsRef = in.SubjectsRef
	rb.PolicyInline = in.PolicyInline
	rb.PolicyRef = in.PolicyRef

	diags := fwtype.RoleBindingTyp.Flatten(ctx, d, rb)
	if diags.HasError() {
		return diags
	}
	return nil
}
