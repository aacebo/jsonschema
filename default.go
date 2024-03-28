package jsonschema

// https://json-schema.org/understanding-json-schema/reference/annotations#annotations
var _default = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		return errs
	},
}
