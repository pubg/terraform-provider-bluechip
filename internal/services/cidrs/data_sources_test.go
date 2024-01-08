package cidrs

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
					resource.TestCheckResourceAttr("data.bluechip_cidrs.current", "id", "office"),
					resource.TestCheckResourceAttrWith("data.bluechip_cidrs.current", "items.#", func(value string) error {
						tflog.Debug(context.Background(), "Search Result", map[string]interface{}{"items": value})
						return nil
					}),
					resource.TestCheckResourceAttrWith("data.bluechip_cidrs.current2", "items.#", func(value string) error {
						tflog.Debug(context.Background(), "Search Result", map[string]interface{}{"items": value})
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourcesConfig = `
data "bluechip_cidrs" "current" {
  namespace = "office"
}

data "bluechip_cidrs" "current2" {
  namespace = "office"

  filter {
	operator = "fuzzy"
	field = "metadata.name"
	value = "console"
  }
}
`
