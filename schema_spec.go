package jsonschema

// https://json-schema.org/understanding-json-schema/reference/schema#declaring-a-dialect
var schemaSpec = Keyword{
	compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.(string)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "$schema",
				Message: `must be a "string"`,
			})
		}

		return errs
	},
}
