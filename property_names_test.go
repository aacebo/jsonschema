package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestPropertyNames(t *testing.T) {
	RunAll("./testcases/object/property_names", jsonschema.New(), t)
}

func BenchmarkPropertyNames(b *testing.B) {
	RunAllBench("./testcases/object/property_names", jsonschema.New(), b)
}
