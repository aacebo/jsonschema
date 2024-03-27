package jsonschema

// https://json-schema.org/understanding-json-schema/reference/conditionals#dependentRequired
var dependencies = Keyword{
	compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		_, ok := ctx.Value.(map[string][]string)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "dependencies",
				Message: `must be a "map[string][]string"`,
			})
		}

		return errs
	},
}
