package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
)

import (
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/accounts"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/applications"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/cidrs"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/clusterrolebindings"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/clusters"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/images"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/namespaces"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/oidcauths"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/rolebindings"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/users"
	_ "github.com/pubg/terraform-provider-bluechip/internal/services/vendors"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	_ string = "dev"
	_ string = "local"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debugMode,
		ProviderAddr: "registry.terraform.io/pubg/bluechip",
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(opts)
}
