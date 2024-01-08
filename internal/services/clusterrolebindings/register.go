package clusterrolebindings

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_clusterrolebinding", NewResource().Resource())
	provider.RegisterDataSource("bluechip_clusterrolebinding", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_clusterrolebindings", NewDataSources().Resource())
}
