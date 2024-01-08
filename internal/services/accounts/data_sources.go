package accounts

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSources[bluechip_models.Account, bluechip_models.AccountSpec]{
		Timeout: 30 * time.Second,
		Gvk:     bluechip_models.AccountGvk,

		FilterType:   fwtype.FilterType{},
		MetadataType: fwservices.NamespacedDataSourcesMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
