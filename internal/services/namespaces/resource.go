package namespaces

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/bluechip_client/bluechip_models"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwservices"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

func NewResource() fwservices.ResourceFactory {
	return &fwservices.ClusterTerraformResource[bluechip_models.Namespace, bluechip_models.EmptySpec]{
		Gvk: bluechip_models.NamespaceGvk,
		Schema: map[string]*schema.Schema{
			"metadata": fwtype.MetadataTyp.Schema(false, false),
		},
		Timeout:       30 * time.Second,
		SpecExpander:  specTyp.Expand,
		SpecFlattener: specTyp.Flatten,
		Constructor: func() bluechip_models.Namespace {
			return bluechip_models.Namespace{
				TypeMeta:          &bluechip_models.TypeMeta{},
				MetadataContainer: &bluechip_models.MetadataContainer{},
			}
		},
	}
}
