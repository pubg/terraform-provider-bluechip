package main

import (
	"flag"

	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

import (
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/accounts"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/applications"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/cidrs"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/clusterrolebindings"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/clusters"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/images"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/namespaces"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/oidcauths"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/rolebindings"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/roles"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/users"
	_ "git.projectbro.com/Devops/terraform-provider-bluechip/internal/services/vendors"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	version string = "dev"
	commit  string = "local"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debugMode,
		ProviderAddr: "registry.terraform.io/pubg/bluechip",
		ProviderFunc: provider.Provider(version, commit),
	}

	plugin.Serve(opts)
}
