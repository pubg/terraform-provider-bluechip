package accounts

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
					resource.TestCheckResourceAttr("bluechip_account.current", "id", "default/test2"),
					resource.TestCheckResourceAttr("bluechip_account.current", "metadata.0.name", "test2"),
					resource.TestCheckResourceAttrWith("bluechip_account.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_account" "current" {
  metadata {
    name = "test2"
    namespace = "default"
  }
  spec {
    account_id = "12398213"
	display_name = "test"
	description = "test"
	alias = "test"
	vendor = "AWS"
	regions = ["test"]
  }
}
`
