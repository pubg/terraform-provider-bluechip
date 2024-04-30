package clusters

import (
	"testing"

	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/testacc"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccResourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("bluechip_cluster.current", "id", "default/test"),
					resource.TestCheckResourceAttr("bluechip_cluster.current", "metadata.0.name", "test"),
					resource.TestCheckResourceAttrWith("bluechip_cluster.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_cluster" "current" {
  metadata {
    name = "test"
    namespace = "default"
  }
  spec {
    project = "pubg"
    environment = "dev"
    organization_unit = "devops"
    platform = "pc"
    pubg {
      infra = "common"
      site = "devops"
    }
    vendor {
      name = "AWS"
      account_id = "12398213"
      engine = "EKS"
      region = "ap-northeast-2"
    }
    kubernetes {
      endpoint = "https://api.devops.dev.pubg.com"
      ca_cert = "-----BEGIN CERTIFI"
      sa_issuer = "https://login.microsoftonline.com/1a27bdbf-e6cc-4e33-85d2-e1c81bad930a/v2.0"
      version = "1.28"
    }
  }
}
`
