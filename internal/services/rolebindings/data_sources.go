package rolebindings

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSources[bluechip.RoleBinding, bluechip.RoleBindingSpec]{
		Gvk:     bluechip.RoleBindingGvk,
		Timeout: 30 * time.Second,

		FilterType:       fwtype.FilterType{},
		MetadataType:     fwservices.NamespacedDataSourcesMetadataType,
		SpecType:         &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
		DebuilderFactory: &DebuilderFactory{},
	}
}
