package bluechip_models

type TypeMeta struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

func (m *TypeMeta) GetApiVersion() string {
	return m.ApiVersion
}

func (m *TypeMeta) SetApiVersion(apiVersion string) {
	m.ApiVersion = apiVersion
}

func (m *TypeMeta) GetKind() string {
	return m.Kind
}

func (m *TypeMeta) SetKind(kind string) {
	m.Kind = kind
}

type MetadataContainer struct {
	Container Metadata `json:"metadata"`
}

func (m *MetadataContainer) GetMetadata() Metadata {
	return m.Container
}

func (m *MetadataContainer) SetMetadata(metadata Metadata) {
	m.Container = metadata
}

type SpecContainer[C BaseSpec] struct {
	Container C `json:"spec"`
}

func (s *SpecContainer[C]) GetSpec() C {
	return s.Container
}

func (s *SpecContainer[C]) SetSpec(spec C) {
	s.Container = spec
}
