package jsonschema

import (
	"fmt"
	"reflect"
	"regexp"
)

// https://json-schema.org/understanding-json-schema/reference/string#regexp
var pattern = Keyword{
	compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		str, ok := ctx.Value.(string)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "pattern",
				Message: `must be a "string"`,
			})

			return errs
		}

		_, err := regexp.Compile(str)

		if err != nil {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "pattern",
				Message: err.Error(),
			})
		}

		return errs
	},
	validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		value := reflect.Indirect(reflect.ValueOf(input))

		if value.Kind() != reflect.String {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "pattern",
				Message: fmt.Sprintf(
					`"%s" should be "string"`,
					value.Kind().String(),
				),
			})

			return errs
		}

		expr := regexp.MustCompile(ctx.Value.(string))

		if !expr.MatchString(value.String()) {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "pattern",
				Message: fmt.Sprintf(
					`"%s" does not match`,
					value.String(),
				),
			})
		}

		return errs
	},
}
