package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestAllOf(t *testing.T) {
	RunAll("./testcases/all_of", jsonschema.New(), t)
}

func BenchmarkAllOf(b *testing.B) {
	RunAllBench("./testcases/all_of", jsonschema.New(), b)
}
