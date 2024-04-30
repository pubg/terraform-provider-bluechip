package clusterrolebindings

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip.ClusterRoleBinding, bluechip.ClusterRoleBindingSpec]{
		Gvk:     bluechip.ClusterRoleBindingGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.ClusterDataSourceMetadataType,
		SpecType:         &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
		DebuilderFactory: &DebuilderFactory{},
	}
}
