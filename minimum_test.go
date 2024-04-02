package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestMinimum(t *testing.T) {
	RunAll("./testcases/number/minimum", jsonschema.New(), t)
	RunAll("./testcases/integer/minimum", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			Integer().
			Minimum(20).
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, 21)

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkMinimum(b *testing.B) {
	RunAllBench("./testcases/number/minimum", jsonschema.New(), b)
	RunAllBench("./testcases/integer/minimum", jsonschema.New(), b)
}
