package fwbuilder

import "git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"

type ResourceBuilderFactory[T api_meta.ApiResource] interface {
	New() ResourceBuilder[T]
}

type ResourceBuilder[T api_meta.ApiResource] interface {
	Set(field string, value any) error
	Build() T
}

const (
	FieldApiVersion = "apiVersion"
	FieldKind       = "kind"
	FieldMetadata   = "metadata"
	FieldSpec       = "spec"
	FieldStatus     = "status"
)
