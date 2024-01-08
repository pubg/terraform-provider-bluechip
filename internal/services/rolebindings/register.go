package rolebindings

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_rolebinding", NewResource().Resource())
	provider.RegisterDataSource("bluechip_rolebinding", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_rolebindings", NewDataSources().Resource())
}
