package applications

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pubg/terraform-provider-bluechip/internal/testacc"
)

func TestWhoamiDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testacc.TestAccPreCheck(t) },
		ProtoV5ProviderFactories: testacc.TestAccProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testacc.CombinedConfig(TestAccWhoAmiDataSourceConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "id", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "name", "admin"),
					resource.TestCheckResourceAttr("data.bluechip_whoami.current", "groups.0", "system-admin"),
				),
			},
		},
	})
}

const TestAccWhoAmiDataSourceConfig = `
data "bluechip_whoami" "current" {
}
`
