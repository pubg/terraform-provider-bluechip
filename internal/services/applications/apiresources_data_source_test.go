package applications

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pubg/terraform-provider-bluechip/internal/testacc"
)

func TestAccApiResourcesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccApiResourcesDataSourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_apiresources.current", "id", "api-resources"),
					resource.TestCheckResourceAttrWith("data.bluechip_apiresources.current", "api_resources.0.group", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccApiResourcesDataSourceConfig = `
data "bluechip_apiresources" "current" {
}
`
