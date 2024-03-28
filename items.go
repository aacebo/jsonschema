package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array#items
func items(key string) Keyword {
	return Keyword{
		Default: Schema{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			switch v := config.Interface().(type) {
			case map[string]any:
				return ns.compile(fmt.Sprintf("%s/%s", ctx.Path, key), v)
			case []any:
				for i, s := range v {
					schema, ok := s.(map[string]any)

					if !ok {
						errs = append(errs, SchemaError{
							Path:    fmt.Sprintf("%s/%s/%d", ctx.Path, key, i),
							Keyword: key,
							Message: `must be a "Schema"`,
						})

						continue
					}

					_errs := ns.compile(fmt.Sprintf("%s/%s/%d", ctx.Path, key, i), schema)

					if len(_errs) > 0 {
						errs = append(errs, _errs...)
					}
				}

				break
			default:
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "Schema" or "[]Schema"`,
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

			items, ok := value.Interface().([]any)

			if !ok {
				return errs
			}

			switch v := config.Interface().(type) {
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
}
