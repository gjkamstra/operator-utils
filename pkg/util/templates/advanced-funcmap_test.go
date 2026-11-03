package templates

import (
	"testing"

	"github.com/go-logr/logr"
)

func TestRequiredPreservesLiteralWarning(t *testing.T) {
	required := AdvancedTemplateFuncMap(nil, logr.Discard())["required"].(func(string, interface{}) (interface{}, error))
	const warning = "value is required: 100% complete %s"

	for _, value := range []interface{}{nil, ""} {
		got, err := required(warning, value)
		if got != value || err == nil || err.Error() != warning {
			t.Errorf("required(%q, %#v) = (%#v, %v)", warning, value, got, err)
		}
	}
}
