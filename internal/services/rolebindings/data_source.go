package rolebindings

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.NamespacedTerraformDataSource[bluechip_models.RoleBinding, bluechip_models.RoleBindingSpec]{
		Gvk:     bluechip_models.RoleBindingGvk,
		Timeout: 30 * time.Second,

		MetadataType: fwservices.NamespacedDataSourceMetadataType,
		SpecType:     &SpecType{RbType: fwtype.RoleBindingType{Computed: true}},
	}
}
