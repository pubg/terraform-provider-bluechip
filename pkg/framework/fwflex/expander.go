package fwflex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Expander[T any] func(ctx context.Context, d *schema.ResourceData, out *T) diag.Diagnostics

func ExtractSingleBlock(ctx context.Context, d *schema.ResourceData, blockName string) (map[string]any, diag.Diagnostics) {
	rawBlocks, ok := d.GetOk(blockName)
	if !ok {
		return nil, diag.Errorf("block %q is required", blockName)
	}

	rawBlockList, ok := rawBlocks.([]any)
	if !ok {
		return nil, diag.Errorf("block %q is not a list", blockName)
	}

	if len(rawBlockList) == 0 || rawBlockList[0] == nil {
		return nil, diag.Errorf("block %q is empty", blockName)
	}

	data, ok := rawBlockList[0].(map[string]any)
	if !ok {
		return nil, diag.Errorf("block %q is not a object", blockName)
	}

	return data, nil
}

func ExpandMap(attr map[string]any) map[string]string {
	m := make(map[string]string)
	for k, v := range attr {
		m[k] = v.(string)
	}
	return m
}

func ExpandStringList(attr []any) []string {
	l := make([]string, len(attr))
	for i, v := range attr {
		l[i] = v.(string)
	}
	return l
}

func ExpandMapList(attr []any) []map[string]any {
	l := make([]map[string]any, len(attr))
	for i, v := range attr {
		l[i] = v.(map[string]any)
	}
	return l
}

func ExpandStringSet(set *schema.Set) []string {
	l := make([]string, set.Len())
	for i, v := range set.List() {
		l[i] = v.(string)
	}
	return l
}

func ExpandStringFromMap(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		return v.(string)
	}
	return ""
}

func ExpandStringPFromMap(m map[string]any, key string) *string {
	if v, ok := m[key]; ok {
		vv := v.(string)
		return &vv
	}
	return nil
}
