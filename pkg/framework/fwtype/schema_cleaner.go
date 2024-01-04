package fwtype

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
