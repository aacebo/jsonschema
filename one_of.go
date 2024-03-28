package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/combining#oneOf
func oneOf(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}
			schemas, ok := config.Interface().([]any)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `should be a "[]Schema"`,
				})

				return errs
			}

			for i, s := range schemas {
				path := fmt.Sprintf("%s/%s/%d", ctx.Path, key, i)
				schema, ok := s.(map[string]any)

				if !ok {
					errs = append(errs, SchemaError{
						Path:    path,
						Keyword: key,
						Message: `should be a "Schema"`,
					})

					continue
				}

				_errs := ns.compile(path, schema)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}
			valid := 0

			for i := 0; i < config.Len(); i++ {
				index := config.Index(i).Elem()

				if index.Kind() != reflect.Map {
					continue
				}

				_errs := ns.validate(
					ctx.Path,
					index.Interface().(map[string]any),
					value.Interface(),
				)

				if len(_errs) == 0 {
					valid++
				}
			}

			if valid != 1 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must match one schema",
				})
			}

			return errs
		},
	}
}
