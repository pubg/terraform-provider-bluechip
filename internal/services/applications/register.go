package applications

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterDataSource("bluechip_whoami", NewWhoamiDataSource())
	provider.RegisterDataSource("bluechip_apiresources", NewApiResourcesDataSource())
}
