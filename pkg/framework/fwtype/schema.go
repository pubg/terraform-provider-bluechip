package fwtype

import (
	"reflect"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CleanForDataSourceMap(s map[string]*schema.Schema) {
	for _, scheme := range s {
		CleanForDataSource(scheme)
	}
}

func CleanForDataSource(scheme *schema.Schema) {
	scheme.ValidateFunc = nil
	scheme.ValidateDiagFunc = nil
	scheme.MinItems = 0
	scheme.MaxItems = 0

	if resource, ok := scheme.Elem.(*schema.Resource); ok {
		CleanForDataSourceMap(resource.Schema)
	}
}

func SetBlock(d *schema.ResourceData, key string, value any) diag.Diagnostics {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return nil
	}

	if err := d.Set(key, []any{value}); err != nil {
		return append(diag.Diagnostics{}, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Failed to set value",
			Detail:        err.Error(),
			AttributePath: cty.GetAttrPath(key),
		})
	}
	return nil
}
