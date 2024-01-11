package fwdiag

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func FromErr(err error, summary string) diag.Diagnostics {
	return diag.Diagnostics{
		{
			Severity: diag.Error,
			Summary:  summary,
			Detail:   err.Error(),
		},
	}
}

func FromErrF(err error, format string, args ...any) diag.Diagnostics {
	return FromErr(err, fmt.Sprintf(format, args...))
}
