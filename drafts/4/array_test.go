package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	test "jsonschema/test"
	"testing"
)

func TestArray(t *testing.T) {
	test.RunAll("./testcases/array", jsonschema.New(), t)
}
