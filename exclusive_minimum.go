package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func exclusiveMinimum(key string) Keyword {
	return Keyword{
		Default: false,
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

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
				if !config.CanFloat() && config.CanConvert(reflect.TypeOf(0.0)) {
					config = config.Convert(reflect.TypeOf(0.0))
				}

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
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))
			value := reflect.Indirect(reflect.ValueOf(input))

			if !value.IsValid() {
				return errs
			}

			if !config.CanFloat() && config.CanConvert(reflect.TypeOf(0.0)) {
				config = config.Convert(reflect.TypeOf(0.0))
			}

			if !value.CanFloat() && value.CanConvert(reflect.TypeOf(0.0)) {
				value = value.Convert(reflect.TypeOf(0.0))
			}

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

				if !minimum.CanFloat() && minimum.CanConvert(reflect.TypeOf(0.0)) {
					minimum = minimum.Convert(reflect.TypeOf(0.0))
				}

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
			} else if config.CanFloat() {
				if !config.CanFloat() && config.CanConvert(reflect.TypeOf(0.0)) {
					config = config.Convert(reflect.TypeOf(0.0))
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
			} else {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "boolean" or "number"`,
				})
			}

			return errs
		},
	}
}
