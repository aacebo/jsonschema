package jsonschema

import "reflect"

// https://json-schema.org/understanding-json-schema/reference/comments#comments
var comment = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		config := reflect.Indirect(reflect.ValueOf(ctx.Value))

		if config.Kind() != reflect.String {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "$comment",
				Message: `must be a "string"`,
			})
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		return errs
	},
}
