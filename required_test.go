package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestRequired(t *testing.T) {
	RunAll("./testcases/object/required", jsonschema.New(), t)
}
