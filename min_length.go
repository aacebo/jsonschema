package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/string#length
func minLength(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			fminLength, ok := ctx.Value.(float64)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be an "int"`,
				})

				return errs
			}

			minLength := int(fminLength)

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
			value := reflect.Indirect(reflect.ValueOf(input))

			if value.Kind() != reflect.String {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%s" should be "string"`,
						value.Kind().String(),
					),
				})

				return errs
			}

			if value.Len() < int(ctx.Value.(float64)) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`length "%d" is less than "%d"`,
						value.Len(),
						int(ctx.Value.(float64)),
					),
				})
			}

			return errs
		},
	}
}
