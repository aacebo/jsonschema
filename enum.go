package jsonschema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/enum#enumerated-values
var enum = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.([]any)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "enum",
				Message: `must be an "array"`,
			})

			return errs
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		value := reflect.Indirect(reflect.ValueOf(input))
		options := ctx.Value.([]any)

		if !value.Comparable() {
			b, _ := json.Marshal(value.Interface())
			value = reflect.ValueOf(string(b))
		}

		for _, o := range options {
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
			Keyword: "enum",
			Message: fmt.Sprintf(
				`must be one of %v`,
				options,
			),
		})

		return errs
	},
}
