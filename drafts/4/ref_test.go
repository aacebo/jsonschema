package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	test "jsonschema/test"
	"testing"
)

func TestRef(t *testing.T) {
	test.RunAll("./testcases/ref", jsonschema.New(), t)
}
