package users

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewDataSource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformDataSource[bluechip_models.User, bluechip_models.UserSpec]{
		Gvk: bluechip_models.UsersGvk,
		Schema: map[string]*schema.Schema{
			"metadata": fwtype.MetadataTyp.Schema(false, true),
			"spec":     specTyp.Schema(true),
		},
		Timeout:       30 * time.Second,
		SpecExpander:  specTyp.Expand,
		SpecFlattener: specTyp.Flatten,
	}
}
