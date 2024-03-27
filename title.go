package jsonschema

// https://json-schema.org/understanding-json-schema/reference/annotations
var title = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.(string)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "title",
				Message: `must be a "string"`,
			})
		}

		return errs
	},
}
