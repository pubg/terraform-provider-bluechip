package bluechip_models

import (
	"fmt"
	"strings"
)

type GroupVersionKind struct {
	Group      string `json:"group"`
	Version    string `json:"version"`
	Kind       string `json:"kind"`
	KindPlural string `json:"kindPlural"`
}

func (g GroupVersionKind) ToApiPath() string {
	return fmt.Sprintf("/apis/%s/%s/%s", g.Group, g.Version, strings.ToLower(g.KindPlural))
}

func (g GroupVersionKind) ToTypeMeta() TypeMeta {
	meta := TypeMeta{}
	meta.SetApiVersion(fmt.Sprintf("%s/%s", g.Group, g.Version))
	meta.SetKind(g.Kind)
	return meta
}

var NamespaceGvk = GroupVersionKind{
	Group:      "core",
	Version:    "v1",
	Kind:       "Namespace",
	KindPlural: "Namespaces",
}

var ClusterRoleBindingGvk = GroupVersionKind{
	Group:      "rbac.bluechip.pubg.io",
	Version:    "v1",
	Kind:       "ClusterRoleBinding",
	KindPlural: "ClusterRoleBindings",
}

var RoleBindingGvk = GroupVersionKind{
	Group:      "rbac.bluechip.pubg.io",
	Version:    "v1",
	Kind:       "RoleBinding",
	KindPlural: "RoleBindings",
}

var VendorGvk = GroupVersionKind{
	Group:      "bluechip.pubg.io",
	Version:    "v1",
	Kind:       "Vendor",
	KindPlural: "Vendors",
}

var ClusterGvk = GroupVersionKind{
	Group:      "bluechip.pubg.io",
	Version:    "v1",
	Kind:       "Cluster",
	KindPlural: "Clusters",
}

var AccountGvk = GroupVersionKind{
	Group:      "bluechip.pubg.io",
	Version:    "v1",
	Kind:       "Account",
	KindPlural: "Accounts",
}

var CidrGvk = GroupVersionKind{
	Group:      "bluechip.pubg.io",
	Version:    "v1",
	Kind:       "Cidr",
	KindPlural: "Cidrs",
}

var ImageGvk = GroupVersionKind{
	Group:      "bluechip.pubg.io",
	Version:    "v1",
	Kind:       "Image",
	KindPlural: "Images",
}

var OidcAuthGvk = GroupVersionKind{
	Group:      "auth.bluechip.pubg.io",
	Version:    "v1",
	Kind:       "OidcAuth",
	KindPlural: "OidcAuths",
}

var UsersGvk = GroupVersionKind{
	Group:      "auth.bluechip.pubg.io",
	Version:    "v1",
	Kind:       "User",
	KindPlural: "Users",
}
