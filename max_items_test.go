package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaxItems(t *testing.T) {
	RunAll("./testcases/array/max_items", jsonschema.New(), t)
}

func BenchmarkMaxItems(b *testing.B) {
	RunAllBench("./testcases/array/max_items", jsonschema.New(), b)
}
