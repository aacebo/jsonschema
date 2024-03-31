package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMinProperties(t *testing.T) {
	RunAll("./testcases/object/min_properties", jsonschema.New(), t)
}
