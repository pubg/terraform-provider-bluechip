package oidcauths

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
					resource.TestCheckResourceAttr("bluechip_oidcauth.current", "id", "my-test"),
					resource.TestCheckResourceAttr("bluechip_oidcauth.current", "metadata.0.name", "my-test"),
					resource.TestCheckResourceAttrSet("bluechip_oidcauth.current", "metadata.0.creation_timestamp"),
					resource.TestCheckResourceAttrWith("bluechip_oidcauth.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_oidcauth" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    username_claim= "sub"
	username_prefix= "string"
	issuer = "https://accounts.google.com/"
	client_id = "string"
	required_claims = ["string"]
	groups_claim = "string"
	groups_prefix = "string"
    attribute_mapping {
      from               = "namespace_path"
      from_path_resolver = "bare"
      to                 = "namespace_path"
    }
    attribute_mapping {
      from               = "project_path"
      from_path_resolver = "bare"
      to                 = "project_path"
    }
    attribute_mapping {
      from               = "pipeline_source"
      from_path_resolver = "bare"
      to                 = "pipeline_source"
    }
    attribute_mapping {
      from               = "ref_path"
      from_path_resolver = "bare"
      to                 = "ref_path"
    }
  }
}
`
