package oidcauths

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

var specTyp = &SpecType{}

func init() {
	provider.RegisterResource("bluechip_oidcauth", NewResource().Resource())
	provider.RegisterDataSource("bluechip_oidcauth", NewDataSource().Resource())
}
