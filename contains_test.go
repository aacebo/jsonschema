package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestContains(t *testing.T) {
	RunAll("./testcases/array/contains", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			Array().
			Contains(jsonschema.Builder().String().Build()).
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, []any{1, "test"})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkContains(b *testing.B) {
	RunAllBench("./testcases/array/contains", jsonschema.New(), b)
}
