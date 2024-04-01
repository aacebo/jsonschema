package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestMinItems(t *testing.T) {
	RunAll("./testcases/array/min_items", jsonschema.New(), t)
}

func BenchmarkMinItems(b *testing.B) {
	RunAllBench("./testcases/array/min_items", jsonschema.New(), b)
}
