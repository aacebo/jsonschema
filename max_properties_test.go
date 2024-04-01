package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaxProperties(t *testing.T) {
	RunAll("./testcases/object/max_properties", jsonschema.New(), t)
}

func BenchmarkMaxProperties(b *testing.B) {
	RunAllBench("./testcases/object/max_properties", jsonschema.New(), b)
}
