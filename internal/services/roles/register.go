package roles

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_role", NewResource().Resource())
	provider.RegisterDataSource("bluechip_role", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_roles", NewDataSources().Resource())
}
