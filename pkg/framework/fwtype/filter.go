package fwtype

import (
	"context"

	"git.projectbro.com/Devops/arcane-client-go/pkg/api_meta"
	"git.projectbro.com/Devops/terraform-provider-bluechip/pkg/framework/fwflex"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var _ TypeHelper[FilterValue] = &FilterType{}

type FilterType struct {
}

func (FilterType) Schema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "Filter is a list of query terms to filter the results by.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"field": {
					Type:        schema.TypeString,
					Description: "Field to use for the query term.",
					Required:    true,
				},
				"operator": {
					Type:        schema.TypeString,
					Description: "Operator to use for the query term. One of ['equals', 'notEquals', 'fuzzy', 'wildcard', 'regex', 'matchPhrase', 'prefix'].",
					Required:    true,
					ValidateFunc: validation.StringInSlice([]string{
						api_meta.OperatorEquals,
						api_meta.OperatorNotEquals,
						api_meta.OperatorFuzzy,
						api_meta.OperatorWildcard,
						api_meta.OperatorRegex,
						api_meta.OperatorMatchPhrase,
						api_meta.OperatorPrefix,
					}, false),
				},
				"value": {
					Type:        schema.TypeString,
					Description: "Value to use for the query term.",
					Required:    true,
				},
			},
		},
	}
}

func (FilterType) Expand(ctx context.Context, d *schema.ResourceData, out *FilterValue) diag.Diagnostics {
	if filter, ok := d.GetOk("filter"); ok {
		for _, filterItem := range fwflex.ExpandMapList(filter.([]any)) {
			out.Filter = append(out.Filter, api_meta.QueryTerm{
				Field:    filterItem["field"].(string),
				Operator: filterItem["operator"].(string),
				Value:    filterItem["value"].(string),
			})
		}
	}

	return nil
}

func (FilterType) Flatten(in FilterValue) map[string]any {
	var filter []map[string]any
	for _, filterItem := range in.Filter {
		filter = append(filter, map[string]any{
			"field":    filterItem.Field,
			"operator": filterItem.Operator,
			"value":    filterItem.Value,
		})
	}
	return map[string]any{
		"filter": filter,
	}
}

func (FilterType) FlattenWithSet(ctx context.Context, d *schema.ResourceData, in FilterValue) diag.Diagnostics {
	var filter []map[string]any
	for _, filterItem := range in.Filter {
		filter = append(filter, map[string]any{
			"field":    filterItem.Field,
			"operator": filterItem.Operator,
			"value":    filterItem.Value,
		})
	}
	if err := d.Set("filter", filter); err != nil {
		return append(diag.FromErr(err), diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Failed to set filter",
			Detail:        err.Error(),
			AttributePath: cty.GetAttrPath("filter"),
		})
	}
	return nil
}

type FilterValue struct {
	Filter []api_meta.QueryTerm
}
