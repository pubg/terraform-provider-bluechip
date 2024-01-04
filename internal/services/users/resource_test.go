package users

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pubg/terraform-provider-bluechip/internal/testacc"
)

func TestAccResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccResourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("bluechip_user.current", "id", "my-test"),
					resource.TestCheckResourceAttr("bluechip_user.current", "metadata.0.name", "my-test"),
					resource.TestCheckResourceAttrSet("bluechip_user.current", "metadata.0.creation_timestamp"),
					resource.TestCheckResourceAttrWith("bluechip_user.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_user" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    password = "tetete"
    groups = ["asdf"]
  }
}
`
