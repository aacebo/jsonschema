package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestContains(t *testing.T) {
	RunAll("./testcases/array/contains", jsonschema.New(), t)
}

func BenchmarkContains(b *testing.B) {
	RunAllBench("./testcases/array/contains", jsonschema.New(), b)
}
