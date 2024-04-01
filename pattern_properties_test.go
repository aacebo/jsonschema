package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestPatternProperties(t *testing.T) {
	RunAll("./testcases/object/pattern_properties", jsonschema.New(), t)
}

func BenchmarkPatternProperties(b *testing.B) {
	RunAllBench("./testcases/object/pattern_properties", jsonschema.New(), b)
}
