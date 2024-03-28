package jsonschema

import "fmt"

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
var additionalItems = Keyword{
	Default: Schema{},
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}

		switch ctx.Value.(type) {
		case bool:
			break
		case map[string]any:
			break
		default:
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "additionalItems",
				Message: `must be a "boolean" or "Schema"`,
			})

			break
		}

		return errs
	},
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		arr, ok := input.([]any)

		if !ok {
			return errs
		}

		items, ok := ctx.Schema["items"].([]any)

		if !ok || len(arr) <= len(items) {
			return errs
		}

		switch v := ctx.Value.(type) {
		case bool:
			if !v {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "additionalItems",
					Message: "too many items",
				})
			}

			break
		case map[string]any:
			for i := len(items); i < len(arr); i++ {
				_errs := ns.validate(fmt.Sprintf("%s/%d", ctx.Path, i), v, arr[i])

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			break
		}

		return errs
	},
}
