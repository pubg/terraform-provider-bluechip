package namespaces

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
					resource.TestCheckResourceAttr("bluechip_namespace.current", "id", "my-test"),
					resource.TestCheckResourceAttr("bluechip_namespace.current", "metadata.0.name", "my-test"),
					resource.TestCheckResourceAttrWith("bluechip_namespace.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_namespace" "current" {
  metadata {
    name = "my-test"
  }
}
`
