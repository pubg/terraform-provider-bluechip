package fwservices

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pubg/terraform-provider-bluechip/pkg/framework/fwtype"
)

type ResourceFactory interface {
	Resource() *schema.Resource
}

var metadataTyp = &fwtype.MetadataType{}

func NamespacedResourceIdentity(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}

func NamespacedResourceIdentityFrom(id string) (string, string) {
	parts := strings.Split(id, "/")
	return parts[0], parts[1]
}

func ClusterResourceIdentity(name string) string {
	return name
}

func ClusterResourceIdentityFrom(id string) string {
	return id
}
