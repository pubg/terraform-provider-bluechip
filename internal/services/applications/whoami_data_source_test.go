package applications

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pubg/terraform-provider-bluechip/internal/testacc"
)

func TestWhoamiDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccWhoAmiDataSourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "id", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "name", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "groups.0", "system-admin"),
				),
			},
		},
	})
}

const TestAccWhoAmiDataSourceConfig = `
data "bluechip_whoami" "current" {
}
`

func TestProviderOidc(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: TestAccProviderOidcConfig + "\n" + TestAccWhoAmiDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "id", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "name", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "groups.0", "system-admin"),
				),
			},
		},
	})
}

const TestAccProviderOidcConfig = `
provider "bluechip" {
  address = "" 
  oidc_auth {
    validator_name = "gitlab"
	token = "asdf"
  }
}
`

func TestProviderAws(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: TestAccProviderAwsConfig + "\n" + TestAccWhoAmiDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "id", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "name", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "groups.0", "system-admin"),
				),
			},
		},
	})
}

const TestAccProviderAwsConfig = `
provider "bluechip" {
  address = ""
  basic_auth {
    username = "admin"
	password = "admin"
  }
  aws_auth {
	cluster_name = "bluechip"
	profile = "pubg-devops2"
	region = "ap-northeast-2"
  }
}
`
