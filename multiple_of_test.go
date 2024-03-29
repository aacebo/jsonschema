package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMultipleOf(t *testing.T) {
	RunAll("./testcases/number/multiple_of", jsonschema.New(), t)
}
