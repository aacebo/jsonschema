package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestDefault(t *testing.T) {
	RunAll("./testcases/default", jsonschema.New(), t)
}
