package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestAdditionalItems(t *testing.T) {
	RunAll("./testcases/array/additional_items", jsonschema.New(), t)
}

func BenchmarkAdditionalItems(b *testing.B) {
	RunAllBench("./testcases/array/additional_items", jsonschema.New(), b)
}
