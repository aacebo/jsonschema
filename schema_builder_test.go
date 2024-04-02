package jsonschema_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestSchemaBuilder(t *testing.T) {
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

		t.Run("should succeed", func(t *testing.T) {
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
