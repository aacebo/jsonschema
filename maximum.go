package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
var maximum = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		maximum, ok := ctx.Value.(float64)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maximum",
				Message: `must be a "float"`,
			})

			return errs
		}

		minimum, ok := ctx.Schema["minimum"].(float64)

		if ok && minimum > maximum {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maximum",
				Message: `must be less than "minimum"`,
			})
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		value := reflect.Indirect(reflect.ValueOf(input))

		if value.Kind() != reflect.Float64 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maximum",
				Message: fmt.Sprintf(
					`"%s" should be "number"`,
					value.Kind().String(),
				),
			})

			return errs
		}

		if value.Float() > ctx.Value.(float64) {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "maximum",
				Message: fmt.Sprintf(
					`"%v" is greater than "%v"`,
					value.Float(),
					ctx.Value.(float64),
				),
			})
		}

		return errs
	},
}
