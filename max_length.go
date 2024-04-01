package jsonschema

import (
	"fmt"
	"reflect"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/string#length
func maxLength(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}
			config = coerce.Int(config)

			if !config.CanInt() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be an "integer"`,
				})

				return errs
			}

			maxLength := config.Int()

			if maxLength < 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be greater than or equal to 0",
				})
			}

			minLength := reflect.ValueOf(ctx.Schema["minLength"])

			if !minLength.IsValid() {
				return errs
			}

			minLength = coerce.Int(minLength)

			if minLength.CanInt() && minLength.Int() > maxLength {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be greater than or equal to "minLength"`,
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || value.Kind() != reflect.String {
				return errs
			}

			config = coerce.Int(config)

			if value.Len() > int(config.Int()) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`length "%d" is greater than "%d"`,
						value.Len(),
						config.Int(),
					),
				})
			}

			return errs
		},
	}
}
