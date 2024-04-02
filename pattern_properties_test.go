package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestPatternProperties(t *testing.T) {
	RunAll("./testcases/object/pattern_properties", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			Object().
			PatternProperties(map[string]jsonschema.Schema{
				"^S_": jsonschema.Builder().String().Build(),
				"^I_": jsonschema.Builder().Integer().Build(),
			}).
			AdditionalProperties(jsonschema.Builder().Number().Build()).
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, struct {
			One   string  `json:"S_one"`
			Two   int     `json:"I_two"`
			Three float64 `json:"I_three"`
		}{"test", 1, 10.0})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkPatternProperties(b *testing.B) {
	RunAllBench("./testcases/object/pattern_properties", jsonschema.New(), b)
}
