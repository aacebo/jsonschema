package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestPattern(t *testing.T) {
	RunAll("./testcases/string/pattern", jsonschema.New(), t)
}

func BenchmarkPattern(b *testing.B) {
	RunAllBench("./testcases/string/pattern", jsonschema.New(), b)
}
