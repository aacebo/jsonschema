package jsonschema

import (
	"fmt"
	"math"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#multiples
var multipleOf = Keyword{
	compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		multipleOf, ok := ctx.Value.(float64)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "multipleOf",
				Message: `must be a "float"`,
			})

			return errs
		}

		if multipleOf < 0 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "multipleOf",
				Message: "must be greater than or equal to 0",
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
				Keyword: "multipleOf",
				Message: fmt.Sprintf(
					`"%s" should be "number"`,
					value.Kind().String(),
				),
			})

			return errs
		}

		if math.Mod(value.Float(), ctx.Value.(float64)) > 0 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "multipleOf",
				Message: fmt.Sprintf(
					`"%v" is not a multiple of "%v"`,
					value.Float(),
					ctx.Value,
				),
			})
		}

		return errs
	},
}
