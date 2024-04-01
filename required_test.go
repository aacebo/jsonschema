package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestRequired(t *testing.T) {
	RunAll("./testcases/object/required", jsonschema.New(), t)
}

func BenchmarkRequired(b *testing.B) {
	RunAllBench("./testcases/object/required", jsonschema.New(), b)
}
