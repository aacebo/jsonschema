package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMinItems(t *testing.T) {
	RunAll("./testcases/array/min_items", jsonschema.New(), t)
}

func BenchmarkMinItems(b *testing.B) {
	RunAllBench("./testcases/array/min_items", jsonschema.New(), b)
}
