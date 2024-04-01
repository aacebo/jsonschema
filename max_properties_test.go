package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestMaxProperties(t *testing.T) {
	RunAll("./testcases/object/max_properties", jsonschema.New(), t)
}

func BenchmarkMaxProperties(b *testing.B) {
	RunAllBench("./testcases/object/max_properties", jsonschema.New(), b)
}
