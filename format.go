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
			format, ok := ctx.Value.(string)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})

				return errs
			}

			if !ns.HasFormat(format) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%s" does not exist`,
						format,
					),
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

			err := ns.Format(ctx.Value.(string), value.String())

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
