package images

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSource[bluechip.Image, bluechip.ImageSpec]{
		Gvk:     bluechip.ImageGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.NamespacedDataSourceMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
