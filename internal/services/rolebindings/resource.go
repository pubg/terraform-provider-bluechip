package rolebindings

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip.RoleBinding, bluechip.RoleBindingSpec]{
		Gvk:     bluechip.RoleBindingGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.NamespacedResourceMetadataType,
		SpecType:         &SpecType{RbType: fwtype.RoleBindingType{Computed: false}},
		DebuilderFactory: &DebuilderFactory{},
		BuilderFactory:   &BuilderFactory{},
	}
}
