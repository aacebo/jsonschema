package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/const#constant-values
func _const(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !config.Comparable() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be a comparable value",
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() {
				return errs
			}

			if value.Kind() != config.Kind() || !value.Equal(config) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`%v does not match %v`,
						value.Interface(),
						config.Interface(),
					),
				})
			}

			return errs
		},
	}
}
