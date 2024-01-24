package roles

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip_models.Role, bluechip_models.RoleSpec]{
		Gvk:     bluechip_models.RoleGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.ClusterDataSourceMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
