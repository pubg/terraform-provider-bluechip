package vendors

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_vendor", NewResource().Resource())
	provider.RegisterDataSource("bluechip_vendor", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_vendors", NewDataSources().Resource())
}
