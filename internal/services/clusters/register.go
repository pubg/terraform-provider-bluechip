package clusters

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_cluster", NewResource().Resource())
	provider.RegisterDataSource("bluechip_cluster", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_clusters", NewDataSources().Resource())
}
