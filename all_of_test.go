package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestAllOf(t *testing.T) {
	RunAll("./testcases/all_of", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			AllOf(
				jsonschema.Builder().
					Object().
					Properties(map[string]jsonschema.Schema{
						"test": jsonschema.Builder().String().Build(),
					}).
					AdditionalProperties(jsonschema.Builder().Integer().Build()).
					Required("test").
					Build(),
			).Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, struct {
			Test  string `json:"test"`
			Other int    `json:"other"`
		}{"test", 1})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkAllOf(b *testing.B) {
	RunAllBench("./testcases/all_of", jsonschema.New(), b)
}
