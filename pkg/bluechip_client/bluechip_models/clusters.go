package bluechip_models

type WhoamiResponse struct {
	BaseSpec `tfsdk:"-"`

	Name       string         `json:"name"`
	Groups     []string       `json:"groups"`
	Attributes map[string]any `json:"attributes"`
}

type ApiResourcesResponse struct {
	BaseSpec `tfsdk:"-"`

	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   any               `json:"metadata"`
	Items      []ApiResourceSpec `json:"items"`
}

type ApiResourceSpec struct {
	BaseSpec

	Group      string `json:"group"`
	Version    string `json:"version"`
	Kind       string `json:"kind"`
	KindPlural string `json:"kindPlural"`
	Namespaced bool   `json:"namespaced"`
}
