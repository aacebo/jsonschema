package jsonschema

// https://json-schema.org/understanding-json-schema/reference/combining#not
func not(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			schema, ok := ctx.Value.(map[string]any)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `should be a "Schema"`,
				})

				return errs
			}

			return ns.compile(ctx.Path, schema)
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			schema, ok := ctx.Value.(map[string]any)

			if !ok {
				return errs
			}

			_errs := ns.validate(ctx.Path, schema, input)

			if len(_errs) == 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "should not match schema",
				})
			}

			return errs
		},
	}
}
