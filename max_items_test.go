package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestMaxItems(t *testing.T) {
	RunAll("./testcases/array/max_items", jsonschema.New(), t)
}

func BenchmarkMaxItems(b *testing.B) {
	RunAllBench("./testcases/array/max_items", jsonschema.New(), b)
}
