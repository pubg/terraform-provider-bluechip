package rolebindings

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSources[bluechip_models.RoleBinding, bluechip_models.RoleBindingSpec]{
		Gvk:     bluechip_models.RoleBindingGvk,
		Timeout: 30 * time.Second,

		FilterType:   fwtype.FilterType{},
		MetadataType: fwservices.NamespacedDataSourcesMetadataType,
		SpecType:     &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
	}
}
