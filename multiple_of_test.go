package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMultipleOf(t *testing.T) {
	RunAll("./testcases/number/multiple_of", jsonschema.New(), t)
	RunAll("./testcases/integer/multiple_of", jsonschema.New(), t)
}

func BenchmarkMultipleOf(b *testing.B) {
	RunAllBench("./testcases/number/multiple_of", jsonschema.New(), b)
	RunAllBench("./testcases/integer/multiple_of", jsonschema.New(), b)
}
