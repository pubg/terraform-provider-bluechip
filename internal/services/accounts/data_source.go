package accounts

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSource[bluechip_models.Account, bluechip_models.AccountSpec]{
		Gvk:     bluechip_models.AccountGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.NamespacedDataSourceMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
