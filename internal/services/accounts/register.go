package accounts

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

var specTyp = &SpecType{}

func init() {
	provider.RegisterResource("bluechip_account", NewResource().Resource())
	provider.RegisterDataSource("bluechip_account", NewDataSource().Resource())
}
