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

			if config.Kind() != reflect.Slice {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `should be a "[]Schema"`,
				})

				return errs
			}

			for i := 0; i < config.Len(); i++ {
				index := config.Index(i).Elem()
				path := fmt.Sprintf("%s/%s/%d", ctx.Path, key, i)

				if index.Kind() != reflect.Map {
					errs = append(errs, SchemaError{
						Path:    path,
						Keyword: key,
						Message: `should be a "Schema"`,
					})

					continue
				}

				_errs := ns.compile(
					ctx.ID,
					path,
					index.Interface().(map[string]any),
				)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.Slice {
				return errs
			}

			for i := 0; i < config.Len(); i++ {
				index := config.Index(i).Elem()

				if index.Kind() != reflect.Map {
					continue
				}

				_errs := ns.validate(
					ctx.ID,
					ctx.Path,
					index.Interface().(map[string]any),
					value.Interface(),
				)

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
