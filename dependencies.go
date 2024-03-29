package jsonschema

import (
	"jsonschema/coerce"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/conditionals#dependentRequired
func dependencies(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}
			config = coerce.MapOf(config, reflect.SliceOf(coerce.STRING_TYPE))

			if config.Kind() != reflect.Map {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "map"`,
				})
			}

			return errs
		},
	}
}
