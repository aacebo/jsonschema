package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestPropertyNames(t *testing.T) {
	RunAll("./testcases/object/property_names", jsonschema.New(), t)
}
