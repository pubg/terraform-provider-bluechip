package users

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pubg/terraform-provider-bluechip/internal/testacc"
)

func TestAccDataSources(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccDataSourcesConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_users.current", "id", "pubg"),
					resource.TestCheckResourceAttrWith("data.bluechip_users.current", "items.#", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourcesConfig = `
data "bluechip_users" "current" {
}
`
