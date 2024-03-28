package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/string#length
func minLength(key string) Keyword {
	return Keyword{
		Default: 0,
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

			if !config.CanInt() && config.CanConvert(reflect.TypeOf(0)) {
				config = config.Convert(reflect.TypeOf(0))
			}

			if !config.CanInt() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be an "integer"`,
				})

				return errs
			}

			minLength := config.Int()

			if minLength < 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be greater than or equal to 0",
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))
			value := reflect.Indirect(reflect.ValueOf(input))

			if value.Kind() != reflect.String {
				return errs
			}

			if !config.CanInt() && config.CanConvert(reflect.TypeOf(0)) {
				config = config.Convert(reflect.TypeOf(0))
			}

			if value.Len() < int(config.Int()) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`length "%d" is less than "%d"`,
						value.Len(),
						config.Int(),
					),
				})
			}

			return errs
		},
	}
}
