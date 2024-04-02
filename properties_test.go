package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestProperties(t *testing.T) {
	RunAll("./testcases/object/properties", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		schema := jsonschema.Builder().
			Object().
			Properties(map[string]jsonschema.Schema{
				"test": jsonschema.Builder().String().Build(),
			}).
			AdditionalProperties(jsonschema.Builder().Integer().Build()).
			Required("test").
			Build()

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

func BenchmarkProperties(b *testing.B) {
	RunAllBench("./testcases/object/properties", jsonschema.New(), b)
}
