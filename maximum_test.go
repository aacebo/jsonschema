package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaximum(t *testing.T) {
	RunAll("./testcases/number/maximum", jsonschema.New(), t)
	RunAll("./testcases/integer/maximum", jsonschema.New(), t)
}

func BenchmarkMaximum(b *testing.B) {
	RunAllBench("./testcases/number/maximum", jsonschema.New(), b)
	RunAllBench("./testcases/integer/maximum", jsonschema.New(), b)
}
