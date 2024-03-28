package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/string#format
func format(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

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
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))
			value := reflect.Indirect(reflect.ValueOf(input))

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
