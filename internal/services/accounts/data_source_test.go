package accounts

import (
	"testing"

	"git.projectbro.com/Devops/terraform-provider-bluechip/internal/testacc"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccDataSourceConfig) + "\n" + TestAccResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_account.current", "id", "default/test2"),
					resource.TestCheckResourceAttr("data.bluechip_account.current", "metadata.0.name", "test2"),
					resource.TestCheckResourceAttrWith("data.bluechip_account.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourceConfig = `
data "bluechip_account" "current" {
  metadata {
    name = "test2"
    namespace = "default"
  }
  depends_on = [bluechip_account.current]
}
`
