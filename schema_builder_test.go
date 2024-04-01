package jsonschema_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestSchemaBuilder(t *testing.T) {
	t.Run("string", func(t *testing.T) {
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
	})

	t.Run("integer", func(t *testing.T) {
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
	})

	t.Run("struct", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			schema := jsonschema.Builder().
				Object().
				Properties(map[string]jsonschema.Schema{
					"test": jsonschema.Builder().String().Build(),
				}).
				AdditionalProperties(jsonschema.Builder().Number().Build()).
				Required("test").
				Build()

			errs := jsonschema.Compile(schema)

			if len(errs) > 0 {
				t.Error(errs)
			}

			errs = jsonschema.Validate(schema, struct {
				Test  string `json:"test"`
				Other string `json:"other"`
			}{"test", "1"})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})

	t.Run("custom", func(t *testing.T) {
		jsonschema.AddKeyword("test", jsonschema.Keyword{
			Compile: func(ns *jsonschema.Namespace, ctx jsonschema.Context, config reflect.Value) []jsonschema.SchemaError {
				errs := []jsonschema.SchemaError{}

				if config.Kind() != reflect.String {
					errs = append(errs, jsonschema.SchemaError{
						Path:    ctx.Path,
						Keyword: "test",
						Message: `must be a "string"`,
					})
				}

				return errs
			},
			Validate: func(ns *jsonschema.Namespace, ctx jsonschema.Context, config, value reflect.Value) []jsonschema.SchemaError {
				errs := []jsonschema.SchemaError{}

				if value.Kind() != reflect.String {
					return errs
				}

				if value.String() != "helloworld" {
					errs = append(errs, jsonschema.SchemaError{
						Path:    ctx.Path,
						Keyword: "test",
						Message: fmt.Sprintf(`must be "%s"`, config.String()),
					})
				}

				return errs
			},
		})

		t.Run("should succed", func(t *testing.T) {
			schema := jsonschema.Builder().
				String().
				Build()

			schema["test"] = "helloworld"
			errs := jsonschema.Compile(schema)

			if len(errs) > 0 {
				t.Error(errs)
			}

			errs = jsonschema.Validate(schema, "helloworld")

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})
}
