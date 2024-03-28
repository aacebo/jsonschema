package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
var exclusiveMaximum = Keyword{
	Default: false,
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.(bool)

		if !ok {
			exclusiveMaximum, ok := ctx.Value.(float64)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMaximum",
					Message: `must be a "bool" or "float"`,
				})

				return errs
			}

			exclusiveMinimum, ok := ctx.Value.(float64)

			if ok && exclusiveMinimum > exclusiveMaximum {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMaximum",
					Message: `must be greater than or equal to "exclusiveMinimum"`,
				})
			}
		} else {
			_, ok := ctx.Schema["maximum"].(float64)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMaximum",
					Message: `"maximum" is required when "bool"`,
				})
			}
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		value := reflect.Indirect(reflect.ValueOf(input))

		if value.Kind() != reflect.Float64 {
			return errs
		}

		switch v := ctx.Value.(type) {
		case bool:
			if !v {
				break
			}

			maximum, _ := ctx.Schema["maximum"].(float64)

			if value.Float() > maximum-1 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMaximum",
					Message: fmt.Sprintf(
						`"%v" is greater than "%v"`,
						value.Float(),
						maximum-1,
					),
				})
			}

			break
		case float64:
			if value.Float() > v {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "exclusiveMaximum",
					Message: fmt.Sprintf(
						`"%v" is greater than "%v"`,
						value.Float(),
						v,
					),
				})
			}

			break
		default:
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "exclusiveMaximum",
				Message: `must be a "bool" or "float"`,
			})

			break
		}

		return errs
	},
}
