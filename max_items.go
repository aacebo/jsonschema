package jsonschema

import (
	"fmt"
	"jsonschema/coerce"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array#length
func maxItems(key string) Keyword {
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

			minItems := reflect.ValueOf(ctx.Schema["minItems"])

			if !minItems.IsValid() {
				return errs
			}

			minItems = coerce.Int(minItems)

			if minItems.CanInt() && minItems.Int() > maxLength {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be greater than or equal to "minItems"`,
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || value.Kind() != reflect.Slice {
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
