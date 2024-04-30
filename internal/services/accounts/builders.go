package accounts

import (
	"fmt"

	"git.projectbro.com/Devops/arcane-client-go/bluechip"
	"git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwbuilder"
)

var _ fwbuilder.ResourceBuilderFactory[bluechip.Account] = &BuilderFactory{}

type BuilderFactory struct {
}

func (f *BuilderFactory) New() fwbuilder.ResourceBuilder[bluechip.Account] {
	return &Builder{obj: bluechip.Account{
		TypeMeta:          &api_meta.TypeMeta{},
		MetadataContainer: &api_meta.MetadataContainer{},
	}}
}

type Builder struct {
	obj bluechip.Account
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
		b.obj.Spec = value.(bluechip.AccountSpec)
	default:
		return fmt.Errorf("unknown field %s", field)
	}
	return nil
}

func (b *Builder) Build() bluechip.Account {
	return b.obj
}

var _ fwbuilder.ResourceDebuilderFactory[bluechip.Account] = &DebuilderFactory{}

type DebuilderFactory struct {
}

func (f *DebuilderFactory) New(obj bluechip.Account) fwbuilder.ResourceDebuilder {
	return &Debuilder{obj: obj}
}

type Debuilder struct {
	obj bluechip.Account
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
		return d.obj.Spec
	default:
		return nil
	}
}
