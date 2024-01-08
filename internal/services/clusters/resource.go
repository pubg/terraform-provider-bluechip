package clusters

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformResource[bluechip_models.Cluster, bluechip_models.ClusterSpec]{
		Gvk:     bluechip_models.ClusterGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.Cluster {
			return bluechip_models.Cluster{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.ClusterSpec]{},
			}
		},

		MetadataType: fwservices.NamespacedResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
