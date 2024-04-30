package images

import "git.projectbro.com/Devops/terraform-provider-bluechip/internal/provider"

func init() {
	provider.RegisterResource("bluechip_image", NewResource().Resource())
	provider.RegisterDataSource("bluechip_image", NewDataSource().Resource())
	provider.RegisterDataSource("bluechip_images", NewDataSources().Resource())
}
