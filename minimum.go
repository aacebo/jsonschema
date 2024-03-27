package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
var minimum = Keyword{
	compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.(float64)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "minimum",
				Message: `must be a "float"`,
			})
		}

		return errs
	},
	validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		value := reflect.Indirect(reflect.ValueOf(input))

		if value.Kind() != reflect.Float64 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "minimum",
				Message: fmt.Sprintf(
					`"%s" should be "number"`,
					value.Kind().String(),
				),
			})

			return errs
		}

		if value.Float() < ctx.Value.(float64) {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "minimum",
				Message: fmt.Sprintf(
					`"%v" is less than "%v"`,
					value.Float(),
					ctx.Value.(float64),
				),
			})
		}

		return errs
	},
}
