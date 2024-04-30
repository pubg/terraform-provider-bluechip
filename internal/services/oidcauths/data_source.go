package oidcauths

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip.OidcAuth, bluechip.OidcAuthSpec]{
		Gvk:     bluechip.OidcAuthGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.ClusterDataSourceMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
