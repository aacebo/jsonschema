package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestFormat(t *testing.T) {
	RunAll("./testcases/string/format", jsonschema.New(), t)
}

func BenchmarkFormat(b *testing.B) {
	RunAllBench("./testcases/string/format", jsonschema.New(), b)
}
