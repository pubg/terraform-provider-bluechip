package roles

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip.Role, bluechip.RoleSpec]{
		Gvk:     bluechip.RoleGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.ClusterDataSourceMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
