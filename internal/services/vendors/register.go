package vendors

import "github.com/pubg/terraform-provider-bluechip/internal/provider"

var specTyp = &SpecType{}

func init() {
	provider.RegisterResource("bluechip_vendor", NewResource().Resource())
	provider.RegisterDataSource("bluechip_vendor", NewDataSource().Resource())
}
