package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	test "jsonschema/test"
	"testing"
)

func TestString(t *testing.T) {
	test.RunAll("./testcases/string", jsonschema.New(), t)
}
