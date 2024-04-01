package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestDefault(t *testing.T) {
	RunAll("./testcases/default", jsonschema.New(), t)
}

func BenchmarkDefault(b *testing.B) {
	RunAllBench("./testcases/default", jsonschema.New(), b)
}
