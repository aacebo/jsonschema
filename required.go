package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/object#required
func required(key string) Keyword {
	return Keyword{
		Default: []string{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.Slice {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be an "array"`,
				})

				return errs
			}

			existing := map[string]bool{}

			for i := 0; i < config.Len(); i++ {
				index := config.Index(i)

				if index.Kind() == reflect.Pointer || index.Kind() == reflect.Interface {
					index = index.Elem()
				}

				if index.Kind() != reflect.String {
					errs = append(errs, SchemaError{
						Path:    fmt.Sprintf("%s/%d", ctx.Path, i),
						Keyword: key,
						Message: `must be a "string"`,
					})

					continue
				}

				if _, ok := existing[index.String()]; ok {
					errs = append(errs, SchemaError{
						Path:    fmt.Sprintf("%s/%d", ctx.Path, i),
						Keyword: key,
						Message: "must be unique",
					})

					continue
				}

				existing[index.String()] = true
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if value.Kind() != reflect.Map {
				return errs
			}

			for i := 0; i < config.Len(); i++ {
				index := config.Index(i).Elem()

				if !value.MapIndex(index).IsValid() || value.MapIndex(index).IsZero() {
					errs = append(errs, SchemaError{
						Path:    fmt.Sprintf("%s/%s", ctx.Path, index.String()),
						Keyword: key,
						Message: `required`,
					})
				}
			}

			return errs
		},
	}
}
