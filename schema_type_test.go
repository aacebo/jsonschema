package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestSchemaType(t *testing.T) {
	RunAll("./testcases/type", jsonschema.New(), t)
}

func BenchmarkSchemaType(b *testing.B) {
	RunAllBench("./testcases/type", jsonschema.New(), b)
}
