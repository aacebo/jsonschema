package jsonschema

import (
	"fmt"
	"jsonschema/coerce"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func maximum(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}
			config = coerce.Float(config)

			if !config.CanFloat() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "number"`,
				})

				return errs
			}

			minimum := reflect.ValueOf(ctx.Schema["minimum"])

			if !minimum.IsValid() {
				return errs
			}

			minimum = coerce.Float(minimum)

			if minimum.CanFloat() && minimum.Float() > config.Float() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be greater than or equal to "minimum"`,
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() {
				return errs
			}

			config = coerce.Float(config)
			value = coerce.Float(value)

			if !value.CanFloat() {
				return errs
			}

			if value.Float() > config.Float() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%v" is greater than "%v"`,
						value.Float(),
						config.Float(),
					),
				})
			}

			return errs
		},
	}
}
