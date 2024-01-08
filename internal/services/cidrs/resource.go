package cidrs

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip_models.Cidr, bluechip_models.CidrSpec]{
		Gvk:     bluechip_models.CidrGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.Cidr {
			return bluechip_models.Cidr{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.CidrSpec]{},
			}
		},

		MetadataType: fwservices.NamespacedResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
