package clusterrolebindings

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip.ClusterRoleBinding, bluechip.ClusterRoleBindingSpec]{
		Gvk:     bluechip.ClusterRoleBindingGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.ClusterResourceMetadataType,
		SpecType:         &SpecType{RbType: fwtype.RoleBindingType{Computed: false}},
		DebuilderFactory: &DebuilderFactory{},
		BuilderFactory:   &BuilderFactory{},
	}
}
