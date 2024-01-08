package users

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip_models.User, bluechip_models.UserSpec]{
		Gvk:     bluechip_models.UsersGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.User {
			return bluechip_models.User{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.UserSpec]{},
			}
		},

		MetadataType: fwservices.ClusterResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
