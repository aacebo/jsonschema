package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestMinLength(t *testing.T) {
	RunAll("./testcases/string/min_length", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			String().
			MinLength(5).
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, "test!")

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkMinLength(b *testing.B) {
	RunAllBench("./testcases/string/min_length", jsonschema.New(), b)
}
