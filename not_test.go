package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestNot(t *testing.T) {
	RunAll("./testcases/not", jsonschema.New(), t)
}

func BenchmarkNot(b *testing.B) {
	RunAllBench("./testcases/not", jsonschema.New(), b)
}
