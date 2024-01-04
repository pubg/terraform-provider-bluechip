package testacc

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joho/godotenv"
	"github.com/pubg/terraform-provider-bluechip/internal/provider"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var TestAccProtoV5ProviderFactories = map[string]func() (tfprotov5.ProviderServer, error){
	"bluechip": func() (tfprotov5.ProviderServer, error) {
		return schema.NewGRPCProviderServer(provider.Provider()), nil
	},
}

func TestAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

func init() {
	var errs []error
	errs = append(errs,
		godotenv.Load(".env"),
		godotenv.Load("../.env"),
		godotenv.Load("../../.env"),
		godotenv.Load("../../../.env"),
	)
	if len(errs) != 4 {
		panic(errs)
	}
}

func ProviderTerraformConfig() string {
	config := `
provider "bluechip" {
  address = "$address" 
  basic_auth {
    username = "$username"
    password = "$password"
  }
}
`
	variables := map[string]string{
		"$address":  os.Getenv("BLUECHIP_ADDRESS"),
		"$username": os.Getenv("BLUECHIP_USERNAME"),
		"$password": os.Getenv("BLUECHIP_PASSWORD"),
	}

	for key, val := range variables {
		config = strings.ReplaceAll(config, key, val)
	}

	return config
}

func CombinedConfig(config ...string) string {
	return ProviderTerraformConfig() + "\n" + strings.Join(config, "\n")
}
