package rolebindings

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSource[bluechip.RoleBinding, bluechip.RoleBindingSpec]{
		Gvk:     bluechip.RoleBindingGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.NamespacedDataSourceMetadataType,
		SpecType:         &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
		DebuilderFactory: &DebuilderFactory{},
	}
}
