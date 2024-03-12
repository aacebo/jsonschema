package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	test "jsonschema/test"
	"testing"
)

func TestNumber(t *testing.T) {
	test.RunAll("./testcases/number", jsonschema.New(), t)
}
