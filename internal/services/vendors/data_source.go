package vendors

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip_models.Vendor, bluechip_models.VendorSpec]{
		Gvk:     bluechip_models.VendorGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.ClusterDataSourceMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
