package namespaces

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip_models.Namespace, bluechip_models.EmptySpec]{
		Gvk:     bluechip_models.NamespaceGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.ClusterDataSourceMetadataType,
		SpecType:     &SpecType{},
	}
}
