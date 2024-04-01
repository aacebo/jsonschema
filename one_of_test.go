package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestOneOf(t *testing.T) {
	RunAll("./testcases/one_of", jsonschema.New(), t)
}

func BenchmarkOneOf(b *testing.B) {
	RunAllBench("./testcases/one_of", jsonschema.New(), b)
}
