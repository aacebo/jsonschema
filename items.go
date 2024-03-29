package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array#items
func items(key string) Keyword {
	return Keyword{
		Default: map[string]any{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			switch config.Kind() {
			case reflect.Map:
				return ns.compile(
					ctx.ID,
					fmt.Sprintf("%s/%s", ctx.Path, key),
					config.Interface().(map[string]any),
				)
			case reflect.Slice:
				for i := 0; i < config.Len(); i++ {
					index := config.Index(i).Elem()

					if index.Kind() != reflect.Map {
						errs = append(errs, SchemaError{
							Path:    fmt.Sprintf("%s/%s/%d", ctx.Path, key, i),
							Keyword: key,
							Message: `must be a "Schema"`,
						})

						continue
					}

					_errs := ns.compile(
						ctx.ID,
						fmt.Sprintf("%s/%s/%d", ctx.Path, key, i),
						index.Interface().(map[string]any),
					)

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

			if !value.IsValid() || value.Kind() != reflect.Slice {
				return errs
			}

			switch config.Kind() {
			case reflect.Map:
				for i := 0; i < value.Len(); i++ {
					index := value.Index(i).Elem()
					_errs := ns.validate(
						ctx.ID,
						fmt.Sprintf("%s/%d", ctx.Path, i),
						config.Interface().(map[string]any),
						index.Interface(),
					)

					if len(_errs) > 0 {
						errs = append(errs, _errs...)
					}
				}

				break
			case reflect.Slice:
				for i := 0; i < config.Len(); i++ {
					index := config.Index(i).Elem()

					if i > value.Len()-1 {
						break
					}

					_errs := ns.validate(
						ctx.ID,
						fmt.Sprintf("%s/%d", ctx.Path, i),
						index.Interface().(map[string]any),
						value.Index(i).Interface(),
					)

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
