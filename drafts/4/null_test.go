package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	test "jsonschema/test"
	"testing"
)

func TestNull(t *testing.T) {
	test.RunAll("./testcases/null", jsonschema.New(), t)
}
