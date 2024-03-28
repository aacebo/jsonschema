package jsonschema

import (
	"fmt"
	"math"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#multiples
func multipleOf(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

			if !config.CanFloat() && config.CanConvert(reflect.TypeOf(0.0)) {
				config = config.Convert(reflect.TypeOf(0.0))
			}

			if !config.CanFloat() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "number"`,
				})

				return errs
			}

			if config.Float() < 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be greater than or equal to 0",
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))
			value := reflect.Indirect(reflect.ValueOf(input))

			if !value.CanFloat() && value.CanConvert(reflect.TypeOf(0.0)) {
				value = value.Convert(reflect.TypeOf(0.0))
			}

			if !value.CanFloat() {
				return errs
			}

			if math.Mod(value.Float(), config.Float()) > 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%v" is not a multiple of "%v"`,
						value.Float(),
						config.Float(),
					),
				})
			}

			return errs
		},
	}
}
