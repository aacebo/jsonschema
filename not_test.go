package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestNot(t *testing.T) {
	RunAll("./testcases/not", jsonschema.New(), t)
}

func BenchmarkNot(b *testing.B) {
	RunAllBench("./testcases/not", jsonschema.New(), b)
}
