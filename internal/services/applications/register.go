package applications

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterDataSource("bluechip_whoami", NewWhoamiDataSource())
}
