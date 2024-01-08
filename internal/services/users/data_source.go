package users

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip_models.User, bluechip_models.UserSpec]{
		Gvk:     bluechip_models.UsersGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.ClusterDataSourceMetadataType,
		SpecType:     &SpecType{Computed: true},
	}
}
