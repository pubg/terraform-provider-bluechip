package namespaces

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pubg/terraform-provider-bluechip/internal/testacc"
)

func TestAccDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccDataSourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_namespace.current", "id", "default"),
					resource.TestCheckResourceAttr("data.bluechip_namespace.current", "metadata.0.name", "default"),
					resource.TestCheckResourceAttrSet("data.bluechip_namespace.current", "metadata.0.creation_timestamp"),
					resource.TestCheckResourceAttrWith("data.bluechip_namespace.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourceConfig = `
data "bluechip_namespace" "current" {
  metadata {
    name = "default"
  }
}
`
