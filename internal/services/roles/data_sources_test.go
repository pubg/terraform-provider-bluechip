package roles

import (
	"testing"

	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/testacc"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSources(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccDataSourcesConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_roles.current", "id", "pubg"),
					resource.TestCheckResourceAttrWith("data.bluechip_roles.current", "items.#", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourcesConfig = `
data "bluechip_roles" "current" {
}
`
