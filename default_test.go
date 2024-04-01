package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestDefault(t *testing.T) {
	RunAll("./testcases/default", jsonschema.New(), t)
}

func BenchmarkDefault(b *testing.B) {
	RunAllBench("./testcases/default", jsonschema.New(), b)
}
