package jsonschema

import (
	"fmt"
	"reflect"
	"regexp"
)

// https://json-schema.org/understanding-json-schema/reference/string#regexp
func pattern(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

			if config.Kind() != reflect.String {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})

				return errs
			}

			_, err := regexp.Compile(config.String())

			if err != nil {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%s" is not a valid regular expression`,
						config.String(),
					),
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))
			value := reflect.Indirect(reflect.ValueOf(input))

			if value.Kind() != reflect.String {
				return errs
			}

			expr := regexp.MustCompile(config.String())

			if !expr.MatchString(value.String()) {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: fmt.Sprintf(
						`"%s" does not match pattern "%s"`,
						value.String(),
						config.String(),
					),
				})
			}

			return errs
		},
	}
}
