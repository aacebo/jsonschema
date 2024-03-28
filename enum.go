package jsonschema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/enum#enumerated-values
func enum(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.Slice {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be an "array"`,
				})

				return errs
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.Comparable() {
				b, _ := json.Marshal(value.Interface())
				value = reflect.ValueOf(string(b))
			}

			for _, o := range config.Interface().([]any) {
				option := reflect.Indirect(reflect.ValueOf(o))

				if !option.Comparable() {
					b, _ := json.Marshal(option.Interface())
					option = reflect.ValueOf(string(b))
				}

				if value.Kind() == option.Kind() {
					if value.Equal(option) {
						return errs
					}
				}
			}

			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: key,
				Message: fmt.Sprintf(
					`must be one of %v`,
					config.Interface(),
				),
			})

			return errs
		},
	}
}
