package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/const#constant-values
func _const(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

			if !config.Comparable() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be a comparable value",
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			value := reflect.Indirect(reflect.ValueOf(input))
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

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
