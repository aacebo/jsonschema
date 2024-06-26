package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestMaxLength(t *testing.T) {
	RunAll("./testcases/string/max_length", jsonschema.New(), t)
}

func BenchmarkMaxLength(b *testing.B) {
	RunAllBench("./testcases/string/max_length", jsonschema.New(), b)
}
