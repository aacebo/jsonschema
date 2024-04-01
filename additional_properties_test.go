package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestAdditionalProperties(t *testing.T) {
	RunAll("./testcases/object/additional_properties", jsonschema.New(), t)
}

func BenchmarkAdditionalProperties(b *testing.B) {
	RunAllBench("./testcases/object/additional_properties", jsonschema.New(), b)
}
