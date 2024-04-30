package cidrs

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip.Cidr, bluechip.CidrSpec]{
		Gvk:     bluechip.CidrGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.NamespacedResourceMetadataType,
		SpecType:         &SpecType{Computed: false},
		DebuilderFactory: &DebuilderFactory{},
		BuilderFactory:   &BuilderFactory{},
	}
}
