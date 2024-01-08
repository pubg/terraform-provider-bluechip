package cidrs

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSource[bluechip_models.Cidr, bluechip_models.CidrSpec]{
		Gvk:     bluechip_models.CidrGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.NamespacedDataSourceMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
