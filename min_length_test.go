package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMinLength(t *testing.T) {
	RunAll("./testcases/string/min_length", jsonschema.New(), t)
}

func BenchmarkMinLength(b *testing.B) {
	RunAllBench("./testcases/string/min_length", jsonschema.New(), b)
}
