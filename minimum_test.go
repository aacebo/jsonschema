package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestSchema(t *testing.T) {
	RunAll("./testcases/number/minimum", jsonschema.New(), t)
}
