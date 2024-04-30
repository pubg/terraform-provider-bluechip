package clusterrolebindings

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
					resource.TestCheckResourceAttr("bluechip_clusterrolebinding.current", "id", "my-test1"),
					resource.TestCheckResourceAttr("bluechip_clusterrolebinding.current", "metadata.0.name", "my-test1"),

					resource.TestCheckResourceAttrSet("bluechip_clusterrolebinding.current", "metadata.0.creation_timestamp"),
					//resource.TestCheckResourceAttrSet("bluechip_clusterrolebinding.policy_inline_path", "metadata.0.creation_timestamp"),
					//resource.TestCheckResourceAttrSet("bluechip_clusterrolebinding.policy_inline_resource", "metadata.0.creation_timestamp"),
					//resource.TestCheckResourceAttrSet("bluechip_clusterrolebinding.current", "spec.0.policy_ref"),
					//resource.TestCheckResourceAttr("bluechip_clusterrolebinding.policy_inline_path", "spec.0.policy_inline.0.actions.0", "read"),
					//resource.TestCheckResourceAttr("bluechip_clusterrolebinding.policy_inline_path", "spec.0.policy_inline.0.paths.0", "/**"),
					//resource.TestCheckResourceAttr("bluechip_clusterrolebinding.policy_inline_resource", "spec.0.policy_inline.0.actions.0", "read"),
					//resource.TestCheckResourceAttr("bluechip_clusterrolebinding.policy_inline_resource", "spec.0.policy_inline.0.resources.0.api_group", "core"),
				),
			},
		},
	})
}

const TestAccResourceConfig = `
resource "bluechip_clusterrolebinding" "current" {
  metadata {
    name = "my-test"
  }
  spec {
    subject_ref {
      kind = "User"
	  name = "my-test"
    }
	role_ref {
	  name = "admin"
	}
  }
}
`
