package fwbuilder

type ResourceDebuilderFactory[T any] interface {
	New(obj T) ResourceDebuilder
}

type ResourceDebuilder interface {
	Get(field string) any
}
