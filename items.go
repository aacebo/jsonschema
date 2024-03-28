package jsonschema

import (
	"fmt"
)

// https://json-schema.org/understanding-json-schema/reference/array#items
var items = Keyword{
	Default: Schema{},
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}

		switch v := ctx.Value.(type) {
		case map[string]any:
			return ns.compile(fmt.Sprintf("%s/items", ctx.Path), v)
		case []any:
			for i, s := range v {
				schema, ok := s.(map[string]any)

				if !ok {
					errs = append(errs, SchemaError{
						Path:    fmt.Sprintf("%s/items/%d", ctx.Path, i),
						Keyword: "items",
						Message: `must be a "Schema"`,
					})

					continue
				}

				_errs := ns.compile(fmt.Sprintf("%s/items/%d", ctx.Path, i), schema)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			break
		default:
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "items",
				Message: `must be a "Schema" or "[]Schema"`,
			})

			break
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		items, ok := input.([]any)

		if !ok {
			return errs
		}

		switch v := ctx.Value.(type) {
		case map[string]any:
			for i, item := range items {
				_errs := ns.validate(fmt.Sprintf("%s/%d", ctx.Path, i), v, item)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			break
		case []any:
			for i, s := range v {
				if i > len(items)-1 {
					break
				}

				schema := s.(map[string]any)
				_errs := ns.validate(fmt.Sprintf("%s/%d", ctx.Path, i), schema, items[i])

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			break
		}

		return errs
	},
}
