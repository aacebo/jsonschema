package jsonschema

import (
	"fmt"
	"jsonschema/coerce"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func exclusiveMaximum(key string) Keyword {
	return Keyword{
		Default: false,
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() == reflect.Bool {
				maximum := reflect.Indirect(reflect.ValueOf(ctx.Schema["maximum"]))

				if !maximum.IsValid() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: `"maximum" is required when "boolean"`,
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

					return errs
				}

				exclusiveMinimum := reflect.Indirect(reflect.ValueOf(ctx.Schema["exclusiveMinimum"]))

				if exclusiveMinimum.CanFloat() && exclusiveMinimum.Float() > config.Float() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: `must be greater than or equal to "exclusiveMinimum"`,
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

				maximum := reflect.Indirect(reflect.ValueOf(ctx.Schema["maximum"]))

				if !maximum.IsValid() {
					return errs
				}

				maximum = coerce.Float(maximum)

				if value.Float() > maximum.Float()-1 {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: fmt.Sprintf(
							`"%v" is greater than "%v"`,
							value.Float(),
							maximum.Float()-1,
						),
					})
				}
			} else {
				config = coerce.Float(config)

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
			}

			return errs
		},
	}
}
