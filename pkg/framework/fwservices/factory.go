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

var NamespacedResourceMetadataType = fwtype.NewMetadataType(true, false)
var NamespacedDataSourceMetadataType = fwtype.NewMetadataType(true, true)
var NamespacedDataSourcesMetadataType = fwtype.NewMetadataType(true, true)

var ClusterResourceMetadataType = fwtype.NewMetadataType(false, false)
var ClusterDataSourceMetadataType = fwtype.NewMetadataType(false, true)
var ClusterDataSourcesMetadataType = fwtype.NewMetadataType(false, true)
