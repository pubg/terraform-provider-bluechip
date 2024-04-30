package oidcauths

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSources() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSources[bluechip.OidcAuth, bluechip.OidcAuthSpec]{
		Gvk:     bluechip.OidcAuthGvk,
		Timeout: 30 * time.Second,

		FilterType:       fwtype.FilterType{},
		MetadataType:     fwservices.ClusterDataSourcesMetadataType,
		SpecType:         &SpecType{Computed: true},
		DebuilderFactory: &DebuilderFactory{},
	}
}
