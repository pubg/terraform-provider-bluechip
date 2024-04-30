package accounts

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSource[bluechip.Account, bluechip.AccountSpec]{
		Gvk:     bluechip.AccountGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.NamespacedDataSourceMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
