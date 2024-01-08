package clusterrolebindings

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip_models.ClusterRoleBinding, bluechip_models.ClusterRoleBindingSpec]{
		Gvk:     bluechip_models.ClusterRoleBindingGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.ClusterDataSourceMetadataType,
		SpecType:     &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
	}
}
