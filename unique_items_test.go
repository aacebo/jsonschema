package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestUniqueItems(t *testing.T) {
	RunAll("./testcases/array/unique_items", jsonschema.New(), t)
}

func BenchmarkUniqueItems(b *testing.B) {
	RunAllBench("./testcases/array/unique_items", jsonschema.New(), b)
}
