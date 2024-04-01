package jsonschema

import (
	"fmt"
	"reflect"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/object#length
func maxProperties(key string) Keyword {
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

			maxProperties := config.Int()

			if maxProperties < 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be greater than or equal to 0",
				})
			}

			minProperties := reflect.ValueOf(ctx.Schema["minProperties"])

			if !minProperties.IsValid() {
				return errs
			}

			minProperties = coerce.Int(minProperties)

			if minProperties.CanInt() && minProperties.Int() > maxProperties {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be greater than or equal to "minProperties"`,
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || value.Kind() != reflect.Map {
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
