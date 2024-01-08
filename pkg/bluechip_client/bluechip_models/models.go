package bluechip_models

type BaseResponse interface {
	_IsBluechipResponse()
}

type BaseSpec interface {
	_IsBluechipSpec()
}

type BaseMetadata interface {
	_IsBluechipMetadata()
}

type ClusterApiResource[Spec BaseSpec] interface {
	BaseResponse

	GetApiVersion() string
	SetApiVersion(apiVersion string)
	GetKind() string
	SetKind(kind string)
	GetMetadata() Metadata
	SetMetadata(m Metadata)
	GetSpec() Spec
	SetSpec(s Spec)
}

type NamespacedApiResource[Spec BaseSpec] interface {
	BaseResponse

	GetApiVersion() string
	SetApiVersion(apiVersion string)
	GetKind() string
	SetKind(kind string)
	GetMetadata() Metadata
	SetMetadata(m Metadata)
	GetSpec() Spec
	SetSpec(s Spec)
}

type ListResponse[Item BaseResponse] interface {
	BaseResponse

	GetApiVersion() string
	SetApiVersion(apiVersion string)
	GetKind() string
	SetKind(kind string)
	GetMetadata() ListMetadata
	GetItems() []Item
}

type EmptySpec struct {
	BaseSpec
}
