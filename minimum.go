package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func minimum(key string) Keyword {
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

			if value.Float() < config.Float() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%v" is less than "%v"`,
						value.Float(),
						config.Float(),
					),
				})
			}

			return errs
		},
	}
}
