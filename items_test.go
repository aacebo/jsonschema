package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestItems(t *testing.T) {
	RunAll("./testcases/array/items", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			Array().
			TupleItems(jsonschema.Builder().String().Build()).
			AdditionalItems(jsonschema.Builder().Integer().Build()).
			MinItems(1).
			MaxItems(2).
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, []any{"test", 1})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkItems(b *testing.B) {
	RunAllBench("./testcases/array/items", jsonschema.New(), b)
}
