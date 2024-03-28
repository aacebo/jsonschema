package jsonschema

import (
	"fmt"
	"jsonschema/coerce"
	"math"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#multiples
func multipleOf(key string) Keyword {
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

			if config.Float() < 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "must be greater than or equal to 0",
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
