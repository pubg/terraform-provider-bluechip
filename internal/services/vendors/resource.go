package vendors

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip_models.Vendor, bluechip_models.VendorSpec]{
		Gvk:     bluechip_models.VendorGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.Vendor {
			return bluechip_models.Vendor{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.VendorSpec]{},
			}
		},

		MetadataType: fwservices.ClusterResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
