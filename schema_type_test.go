package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestSchemaType(t *testing.T) {
	RunAll("./testcases/type", jsonschema.New(), t)
}
