package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestProperties(t *testing.T) {
	RunAll("./testcases/object/properties", jsonschema.New(), t)
}