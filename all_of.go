package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/combining#allOf
func allOf(key string) Keyword {
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
			schemas, ok := config.Interface().([]any)

			if !ok {
				return errs
			}

			for _, s := range schemas {
				schema, ok := s.(map[string]any)

				if !ok {
					continue
				}

				_errs := ns.validate(ctx.Path, schema, value.Interface())

				if len(_errs) > 0 {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: `must match all schemas`,
					})

					return errs
				}
			}

			return errs
		},
	}
}
