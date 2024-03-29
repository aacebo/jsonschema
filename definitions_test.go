package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestDefinitions(t *testing.T) {
	RunAll("./testcases/definitions", jsonschema.New(), t)
}
