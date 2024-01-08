package users

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_user", NewResource().Resource())
	provider.RegisterDataSource("bluechip_user", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_users", NewDataSources().Resource())
}
