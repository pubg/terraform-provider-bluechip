package namespaces

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip_models.Namespace, bluechip_models.EmptySpec]{
		Gvk:     bluechip_models.NamespaceGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.Namespace {
			return bluechip_models.Namespace{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
			}
		},

		MetadataType: fwservices.ClusterResourceMetadataType,
		SpecType:     &SpecType{},
	}
}
