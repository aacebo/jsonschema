package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestAdditionalItems(t *testing.T) {
	RunAll("./testcases/array/additional_items", jsonschema.New(), t)
}

func BenchmarkAdditionalItems(b *testing.B) {
	RunAllBench("./testcases/array/additional_items", jsonschema.New(), b)
}
