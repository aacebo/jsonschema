package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
func additionalItems(key string) Keyword {
	return Keyword{
		Default: Schema{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			switch config.Interface().(type) {
			case bool:
				break
			case map[string]any:
				break
			default:
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "boolean" or "Schema"`,
				})

				break
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() {
				return errs
			}

			arr, ok := value.Interface().([]any)

			if !ok {
				return errs
			}

			items, ok := ctx.Schema["items"].([]any)

			if !ok || len(arr) <= len(items) {
				return errs
			}

			switch v := config.Interface().(type) {
			case bool:
				if !v {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
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
}
