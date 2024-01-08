package rolebindings

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip_models.RoleBinding, bluechip_models.RoleBindingSpec]{
		Gvk:     bluechip_models.RoleBindingGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.RoleBinding {
			return bluechip_models.RoleBinding{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.RoleBindingSpec]{},
			}
		},

		MetadataType: fwservices.NamespacedResourceMetadataType,
		SpecType:     &SpecType{RbType: fwtype.RoleBindingType{Computed: false}},
	}
}
