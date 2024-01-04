package cidrs

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

var specTyp = &SpecType{}

func init() {
	provider.RegisterResource("bluechip_cidr", NewResource().Resource())
	provider.RegisterDataSource("bluechip_cidr", NewDataSource().Resource())
}
