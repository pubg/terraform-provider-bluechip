package vendors

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
					resource.TestCheckResourceAttr("bluechip_vendor.current", "id", "my-test"),
					resource.TestCheckResourceAttr("bluechip_vendor.current", "metadata.0.name", "my-test"),
					resource.TestCheckResourceAttrSet("bluechip_vendor.current", "metadata.0.creation_timestamp"),
					resource.TestCheckResourceAttrWith("bluechip_vendor.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_vendor" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    display_name = "asdf"
	code_name = "AWS"
    short_name = "aws"
    regions = ["asdf"]
  }
}
`
