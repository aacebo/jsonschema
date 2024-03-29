package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestItems(t *testing.T) {
	RunAll("./testcases/array/items", jsonschema.New(), t)
}
