package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestConst(t *testing.T) {
	RunAll("./testcases/const", jsonschema.New(), t)
}

func BenchmarkConst(b *testing.B) {
	RunAllBench("./testcases/const", jsonschema.New(), b)
}
