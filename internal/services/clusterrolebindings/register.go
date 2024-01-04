package clusterrolebindings

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

var specTyp = &SpecType{}

func init() {
	provider.RegisterResource("bluechip_clusterrolebinding", NewResource().Resource())
	provider.RegisterDataSource("bluechip_clusterrolebinding", NewDataSource().Resource())
}
