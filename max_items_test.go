package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaxItems(t *testing.T) {
	RunAll("./testcases/array/max_items", jsonschema.New(), t)
}
