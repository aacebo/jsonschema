package jsonschema

import "fmt"

// https://json-schema.org/understanding-json-schema/reference/combining#oneOf
var oneOf = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}
		schemas, ok := ctx.Value.([]any)

		if !ok {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "oneOf",
				Message: `should be a "[]Schema"`,
			})

			return errs
		}

		for i, s := range schemas {
			path := fmt.Sprintf("%s/oneOf/%d", ctx.Path, i)
			schema, ok := s.(map[string]any)

			if !ok {
				errs = append(errs, SchemaError{
					Path:    path,
					Keyword: "oneOf",
					Message: `should be a "Schema"`,
				})

				continue
			}

			_errs := ns.compile(path, schema)

			if len(_errs) > 0 {
				errs = append(errs, _errs...)
			}
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		schemas, ok := ctx.Value.([]any)

		if !ok {
			return errs
		}

		valid := 0

		for _, s := range schemas {
			schema, ok := s.(map[string]any)

			if !ok {
				continue
			}

			_errs := ns.validate(ctx.Path, schema, input)

			if len(_errs) == 0 {
				valid++
			}
		}

		if valid != 1 {
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "oneOf",
				Message: "must match one schema",
			})
		}

		return errs
	},
}
