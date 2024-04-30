package clusters

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
					resource.TestCheckResourceAttr("data.bluechip_cluster.current", "id", "default/test"),
					resource.TestCheckResourceAttr("data.bluechip_cluster.current", "metadata.0.name", "test"),
					resource.TestCheckResourceAttrWith("data.bluechip_cluster.current", "metadata.0.name", func(value string) error {
						return nil
					}),
				),
			},
		},
	})
}

const TestAccDataSourceConfig = `
data "bluechip_cluster" "current" {
  metadata {
    name = "test"
    namespace = "default"
  }
  depends_on = [bluechip_cluster.current]
}
`
