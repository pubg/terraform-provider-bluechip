package oidcauths

import (
	"time"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip.OidcAuth, bluechip.OidcAuthSpec]{
		Gvk:     bluechip.OidcAuthGvk,
		Timeout: 30 * time.Second,

		MetadataType:     fwservices.ClusterResourceMetadataType,
		SpecType:         &SpecType{Computed: false},
		DebuilderFactory: &DebuilderFactory{},
		BuilderFactory:   &BuilderFactory{},
	}
}
