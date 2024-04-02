package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestPropertyNames(t *testing.T) {
	RunAll("./testcases/object/property_names", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			Object().
			PropertyNames(jsonschema.Builder().Pattern("^S_").Build()).
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, struct {
			One   string  `json:"S_one"`
			Two   int     `json:"S_two"`
			Three float64 `json:"S_three"`
		}{"test", 1, 10.0})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkPropertyNames(b *testing.B) {
	RunAllBench("./testcases/object/property_names", jsonschema.New(), b)
}
