package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestPattern(t *testing.T) {
	RunAll("./testcases/string/pattern", jsonschema.New(), t)
}
