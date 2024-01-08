package oidcauths

import (
	"time"

	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip_models.OidcAuth, bluechip_models.OidcAuthSpec]{
		Gvk:     bluechip_models.OidcAuthGvk,
		Timeout: 30 * time.Second,
		Constructor: func() bluechip_models.OidcAuth {
			return bluechip_models.OidcAuth{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
				SpecContainer:     &bluechip_models.SpecContainer[bluechip_models.OidcAuthSpec]{},
			}
		},

		MetadataType: fwservices.ClusterResourceMetadataType,
		SpecType:     &SpecType{Computed: false},
	}
}
