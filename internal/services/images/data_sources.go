package images

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSources[bluechip.Image, bluechip.ImageSpec]{
		Gvk:     bluechip.ImageGvk,
		Timeout: 30 * time.Second,

		FilterType:       fwtype.FilterType{},
		MetadataType:     fwservices.NamespacedDataSourcesMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
