package cidrs

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_cidr", NewResource().Resource())
	provider.RegisterDataSource("bluechip_cidr", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_cidrs", NewDataSources().Resource())
}
