package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/string#format
func format(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.String {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})

				return errs
			}

			if !ns.HasFormat(config.String()) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%s" does not exist`,
						config.String(),
					),
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() {
				return errs
			}

			if value.Kind() != reflect.String {
				return errs
			}

			err := ns.Format(config.String(), value.String())

			if err != nil {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: err.Error(),
				})
			}

			return errs
		},
	}
}
