package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
var exclusiveMinimum = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.(bool)

		if !ok {
			_, ok = ctx.Value.(float64)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMinimum",
					Message: `must be a "bool" or "float"`,
				})
			}
		} else {
			_, ok := ctx.Schema["minimum"].(float64)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMinimum",
					Message: `"minimum" is required when "bool"`,
				})
			}
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		value := reflect.Indirect(reflect.ValueOf(input))

		if value.Kind() != reflect.Float64 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "exclusiveMinimum",
				Message: fmt.Sprintf(
					`"%s" should be "float"`,
					value.Kind().String(),
				),
			})

			return errs
		}

		switch v := ctx.Value.(type) {
		case bool:
			if !v {
				break
			}

			minimum, _ := ctx.Schema["minimum"].(float64)

			if value.Float() < minimum+1 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMinimum",
					Message: fmt.Sprintf(
						`"%v" is less than "%v"`,
						value.Float(),
						minimum+1,
					),
				})
			}

			break
		case float64:
			if value.Float() < v {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMinimum",
					Message: fmt.Sprintf(
						`"%v" is less than "%v"`,
						value.Float(),
						v,
					),
				})
			}

			break
		default:
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "exclusiveMinimum",
				Message: `must be a "bool" or "float"`,
			})

			break
		}

		return errs
	},
}
