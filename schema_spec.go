package jsonschema

// https://json-schema.org/understanding-json-schema/reference/schema#declaring-a-dialect
func schemaSpec(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			_, ok := ctx.Value.(string)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})
			}

			return errs
		},
	}
}
