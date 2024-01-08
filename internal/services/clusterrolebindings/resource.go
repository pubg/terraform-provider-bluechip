package clusterrolebindings

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip_models.ClusterRoleBinding, bluechip_models.ClusterRoleBindingSpec]{
		Gvk:     bluechip_models.ClusterRoleBindingGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.ClusterRoleBinding {
			return bluechip_models.ClusterRoleBinding{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.ClusterRoleBindingSpec]{},
			}
		},

		MetadataType: fwservices.ClusterResourceMetadataType,
		SpecType:     &SpecType{RbType: fwtype.RoleBindingType{Computed: false}},
	}
}
