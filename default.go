package jsonschema

import "reflect"

// https://json-schema.org/understanding-json-schema/reference/annotations#annotations
func _default(_ string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}
			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}
			return errs
		},
	}
}
