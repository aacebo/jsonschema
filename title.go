package jsonschema

import "reflect"

// https://json-schema.org/understanding-json-schema/reference/annotations
func title(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.String {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})
			}

			return errs
		},
	}
}
