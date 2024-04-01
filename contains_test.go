package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestContains(t *testing.T) {
	RunAll("./testcases/array/contains", jsonschema.New(), t)
}

func BenchmarkContains(b *testing.B) {
	RunAllBench("./testcases/array/contains", jsonschema.New(), b)
}
