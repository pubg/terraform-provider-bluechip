package fwlog

func Field(key string, value any) map[string]any {
	return map[string]any{
		key: value,
	}
}
