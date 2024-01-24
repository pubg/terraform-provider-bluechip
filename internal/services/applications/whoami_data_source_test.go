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
  address = "https://bluechip.example.io"
  auth_flow {
    oidc {
      validator_name = "kubernetes-centre"
      token          = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }
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
  address = "https://bluechip.example.io"
  auth_flow {
    aws {
      cluster_name = "bluechip"
      region       = "us-east-1"
    }
  }
}
`

func TestProviderAwsAutoDiscovery(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: TestAccProviderAwsAutoDiscoveryConfig + "\n" + TestAccWhoAmiDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "id", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "name", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "groups.0", "system-admin"),
				),
			},
		},
	})
}

const TestAccProviderAwsAutoDiscoveryConfig = `
provider "bluechip" {
  address = "https://bluechip.example.io"
  auth_flow {
    aws {
      region = "us-east-1"
	  # profile = "pubg"
    }
  }
}
`

func TestProviderChain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: TestAccProviderChainConfig + "\n" + TestAccWhoAmiDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "id", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "name", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "groups.0", "system:admin"),
				),
			},
		},
	})
}

const TestAccProviderChainConfig = `
provider "bluechip" {
  address = "http://localhost:3000"
  auth_flow {
    basic {
      username = "admin"
      password = "ulizzang"
    }
  }
  auth_flow {
    oidc {
      validator_name = "kubernetes-centre"
      token          = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    }
  }
  auth_flow {
    aws {
      cluster_name = "bluechip2"
      region       = "us-east-1"
    }
  }
}
`
