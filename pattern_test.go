package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestPattern(t *testing.T) {
	RunAll("./testcases/string/pattern", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			String().
			Pattern("^[0-9]*$").
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, "123")

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkPattern(b *testing.B) {
	RunAllBench("./testcases/string/pattern", jsonschema.New(), b)
}
