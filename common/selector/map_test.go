package selector_test

import (
	"testing"

	"github.com/roccowong95/airtable-cli/common/selector"
)

func TestSelectMap(t *testing.T) {
	selector.SelectFromMap(map[string]map[string]interface{}{
		"1": map[string]interface{}{"k": "v"},
	})
}
