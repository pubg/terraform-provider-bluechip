package namespaces

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_namespace", NewResource().Resource())
	provider.RegisterDataSource("bluechip_namespace", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_namespaces", NewDataSources().Resource())
}
