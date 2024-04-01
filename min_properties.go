package jsonschema

import (
	"fmt"
	"reflect"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/object#length
func minProperties(key string) Keyword {
	return Keyword{
		Default: 0,
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

			if config.Int() < 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be greater than or equal to 0",
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
