package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestOneOf(t *testing.T) {
	RunAll("./testcases/one_of", jsonschema.New(), t)
}