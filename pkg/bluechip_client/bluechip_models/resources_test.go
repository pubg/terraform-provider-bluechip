package bluechip_models

import (
	"encoding/json"
	"testing"

	"github.com/yudai/gojsondiff"
)

var jsonResource = `
{
  "apiVersion": "core/v1",
  "kind": "Namespace",
  "metadata": {
    "name": "kube-system",
    "labels": {
      "name": "kube-system"
    },
    "annotations": {
      "scheduler.alpha.kubernetes.io/critical-pod": ""
    }
  }
}
`
var expectedResource = Namespace{
	TypeMeta: &TypeMeta{
		ApiVersion: "core/v1",
		Kind:       "Namespace",
	},
	MetadataContainer: &MetadataContainer{
		Container: Metadata{
			Name: "kube-system",
			Labels: map[string]string{
				"name": "kube-system",
			},
			Annotations: map[string]string{
				"scheduler.alpha.kubernetes.io/critical-pod": "",
			},
		},
	},
}
var differ = gojsondiff.New()

func TestNamespaceMarshal(t *testing.T) {
	buf, err := json.Marshal(expectedResource)
	if err != nil {
		t.Error(err)
		return
	}

	diff, err := differ.Compare([]byte(jsonResource), buf)
	if err != nil {
		t.Error(err)
		return
	}

	if diff.Modified() {
		t.Errorf("Expected equals, but got some diffs %s", diff)
	}
}

func TestNamespaceUnmarshal(t *testing.T) {
	var resource Namespace
	if err := json.Unmarshal([]byte(jsonResource), &resource); err != nil {
		t.Error(err)
		return
	}

	diff := differ.CompareArrays([]any{expectedResource}, []any{resource})
	if diff.Modified() {
		t.Errorf("Expected equals, but got some diffs %s", diff)
	}
}
