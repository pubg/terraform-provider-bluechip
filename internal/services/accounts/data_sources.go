package accounts

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSources[bluechip.Account, bluechip.AccountSpec]{
		Timeout: 30 * time.Second,
		Gvk:     bluechip.AccountGvk,

		FilterType:       fwtype.FilterType{},
		MetadataType:     fwservices.NamespacedDataSourcesMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
