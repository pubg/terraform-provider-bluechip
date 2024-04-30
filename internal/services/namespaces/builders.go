package namespaces

import (
	"fmt"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwbuilder"
)

var _ fwbuilder.ResourceBuilderFactory[bluechip.Namespace] = &BuilderFactory{}

type BuilderFactory struct {
}

func (f *BuilderFactory) New() fwbuilder.ResourceBuilder[bluechip.Namespace] {
	return &Builder{obj: bluechip.Namespace{
		TypeMeta:          &api_meta.TypeMeta{},
		MetadataContainer: &api_meta.MetadataContainer{},
	}}
}

type Builder struct {
	obj bluechip.Namespace
}

func (b *Builder) Set(field string, value any) error {
	switch field {
	case fwbuilder.FieldApiVersion:
		b.obj.ApiVersion = value.(string)
	case fwbuilder.FieldKind:
		b.obj.Kind = value.(string)
	case fwbuilder.FieldMetadata:
		b.obj.SetMetadata(value.(api_meta.Metadata))
	case fwbuilder.FieldSpec:
	default:
		return fmt.Errorf("unknown field %s", field)
	}
	return nil
}

func (b *Builder) Build() bluechip.Namespace {
	return b.obj
}

var _ fwbuilder.ResourceDebuilderFactory[bluechip.Namespace] = &DebuilderFactory{}

type DebuilderFactory struct {
}

func (f *DebuilderFactory) New(obj bluechip.Namespace) fwbuilder.ResourceDebuilder {
	return &Debuilder{obj: obj}
}

type Debuilder struct {
	obj bluechip.Namespace
}

func (d *Debuilder) Get(field string) any {
	switch field {
	case fwbuilder.FieldApiVersion:
		return d.obj.ApiVersion
	case fwbuilder.FieldKind:
		return d.obj.Kind
	case fwbuilder.FieldMetadata:
		return d.obj.GetMetadata()
	case fwbuilder.FieldSpec:
		return nil
	default:
		return nil
	}
}
