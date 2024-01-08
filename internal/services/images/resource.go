package images

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip_models.Image, bluechip_models.ImageSpec]{
		Gvk:     bluechip_models.ImageGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.Image {
			return bluechip_models.Image{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.ImageSpec]{},
			}
		},

		MetadataType: fwservices.NamespacedResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
