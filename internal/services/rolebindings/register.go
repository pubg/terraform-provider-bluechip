package rolebindings

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_rolebinding", NewResource().Resource())
	provider.RegisterDataSource("bluechip_rolebinding", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_rolebindings", NewDataSources().Resource())
}
