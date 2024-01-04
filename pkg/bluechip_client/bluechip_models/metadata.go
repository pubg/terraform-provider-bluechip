package bluechip_models

type Metadata struct {
	BaseMetadata `json:"-"`

	Name              string            `json:"name"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	UpdateTimestamp   string            `json:"UpdateTimestamp,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
}
