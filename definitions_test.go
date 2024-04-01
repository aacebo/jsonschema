package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestDefinitions(t *testing.T) {
	RunAll("./testcases/definitions", jsonschema.New(), t)
}

func BenchmarkDefinitions(b *testing.B) {
	RunAllBench("./testcases/definitions", jsonschema.New(), b)
}
