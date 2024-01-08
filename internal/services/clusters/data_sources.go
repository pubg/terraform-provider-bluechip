package clusters

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSources[bluechip_models.Cluster, bluechip_models.ClusterSpec]{
		Gvk:     bluechip_models.ClusterGvk,
		Timeout: 30 * time.Second,

		FilterType:   fwtype.FilterType{},
		MetadataType: fwservices.NamespacedDataSourcesMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
