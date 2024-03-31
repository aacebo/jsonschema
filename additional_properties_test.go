package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestAdditionalProperties(t *testing.T) {
	RunAll("./testcases/object/additional_properties", jsonschema.New(), t)
}
