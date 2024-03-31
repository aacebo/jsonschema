package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaxProperties(t *testing.T) {
	RunAll("./testcases/object/max_properties", jsonschema.New(), t)
}
