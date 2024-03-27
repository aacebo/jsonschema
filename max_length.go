package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/string#length
var maxLength = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		maxLength, ok := ctx.Value.(int)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maxLength",
				Message: `must be an "int"`,
			})

			return errs
		}

		if maxLength < 0 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maxLength",
				Message: "must be greater than or equal to 0",
			})
		}

		minLength, ok := ctx.Schema["minLength"].(int)

		if ok && minLength > maxLength {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maxLength",
				Message: `must be greater than "minLength"`,
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
				Keyword: "maxLength",
				Message: fmt.Sprintf(
					`"%s" should be "string"`,
					value.Kind().String(),
				),
			})

			return errs
		}

		if value.Len() > ctx.Value.(int) {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maxLength",
				Message: fmt.Sprintf(
					`length "%d" is greater than "%d"`,
					value.Len(),
					ctx.Value.(int),
				),
			})
		}

		return errs
	},
}
