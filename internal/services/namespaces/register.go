package namespaces

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

var specTyp = &SpecType{}

func init() {
	provider.RegisterResource("bluechip_namespace", NewResource().Resource())
	provider.RegisterDataSource("bluechip_namespace", NewDataSource().Resource())
}
