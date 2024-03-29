package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestAnyOf(t *testing.T) {
	RunAll("./testcases/any_of", jsonschema.New(), t)
}
