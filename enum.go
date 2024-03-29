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

			if !value.IsValid() {
				return errs
			}

			if !value.Comparable() {
				b, _ := json.Marshal(value.Interface())
				value = reflect.ValueOf(string(b))
			}

			for i := 0; i < config.Len(); i++ {
				index := config.Index(i).Elem()

				if !index.Comparable() {
					b, _ := json.Marshal(index.Interface())
					index = reflect.ValueOf(string(b))
				}

				if value.Kind() == index.Kind() {
					if value.Equal(index) {
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
