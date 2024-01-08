package clusterrolebindings

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSources[bluechip_models.ClusterRoleBinding, bluechip_models.ClusterRoleBindingSpec]{
		Gvk:     bluechip_models.ClusterRoleBindingGvk,
		Timeout: 30 * time.Second,

		FilterType:   fwtype.FilterType{},
		MetadataType: fwservices.ClusterDataSourcesMetadataType,
		SpecType:     &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
	}
}
