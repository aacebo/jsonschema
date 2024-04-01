package jsonschema

import (
	"fmt"
	"reflect"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func exclusiveMinimum(key string) Keyword {
	return Keyword{
		Default: false,
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() == reflect.Bool {
				minimum := reflect.Indirect(reflect.ValueOf(ctx.Schema["minimum"]))

				if !minimum.IsValid() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: `"minimum" is required when "boolean"`,
					})
				}
			} else {
				config = coerce.Float(config)

				if !config.CanFloat() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: `must be a "boolean" or "number"`,
					})
				}
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() {
				return errs
			}

			value = coerce.Float(value)

			if !value.CanFloat() {
				return errs
			}

			if config.Kind() == reflect.Bool {
				if !config.Bool() {
					return errs
				}

				minimum := reflect.Indirect(reflect.ValueOf(ctx.Schema["minimum"]))

				if !minimum.IsValid() {
					return errs
				}

				minimum = coerce.Float(minimum)

				if value.Float() < minimum.Float()+1 {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: fmt.Sprintf(
							`"%v" is less than "%v"`,
							value.Float(),
							minimum.Float()+1,
						),
					})
				}
			} else {
				config = coerce.Float(config)

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
			}

			return errs
		},
	}
}
