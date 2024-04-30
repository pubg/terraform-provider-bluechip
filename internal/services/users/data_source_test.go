package users

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
				Config: testacc.CombinedConfig(TestAccResourceConfig, TestAccDataSourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_user.current", "id", "my-test"),
					resource.TestCheckResourceAttr("data.bluechip_user.current", "metadata.0.name", "my-test"),
					resource.TestCheckResourceAttrSet("data.bluechip_user.current", "metadata.0.creation_timestamp"),
					resource.TestCheckResourceAttrWith("data.bluechip_user.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourceConfig = `
data "bluechip_user" "current" {
  metadata {
    name = "my-test"
  }
  depends_on = [bluechip_user.current]
}
`
