package jsonschema

// https://json-schema.org/understanding-json-schema/reference/annotations#annotations
func _default(_ string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			return errs
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			return errs
		},
	}
}
