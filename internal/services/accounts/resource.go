package accounts

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip_models.Account, bluechip_models.AccountSpec]{
		Timeout: 30 * time.Second,
		Gvk:     bluechip_models.AccountGvk,
		Constructor: func() bluechip_models.Account {
			return bluechip_models.Account{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.AccountSpec]{},
			}
		},

		MetadataType: fwservices.NamespacedResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
