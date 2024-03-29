package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
func additionalItems(key string) Keyword {
	return Keyword{
		Default: map[string]any{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			switch config.Kind() {
			case reflect.Bool:
				break
			case reflect.Map:
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

			if !value.IsValid() || value.Kind() != reflect.Slice {
				return errs
			}

			items := reflect.Indirect(reflect.ValueOf(ctx.Schema["items"]))

			if !items.IsValid() || items.Kind() != reflect.Slice {
				return errs
			}

			if value.Len() <= items.Len() {
				return errs
			}

			switch config.Kind() {
			case reflect.Bool:
				if !config.Bool() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: "too many items",
					})
				}

				break
			case reflect.Map:
				for i := items.Len(); i < value.Len(); i++ {
					index := value.Index(i)
					_errs := ns.validate(
						fmt.Sprintf("%s/%d", ctx.Path, i),
						config.Interface().(map[string]any),
						index.Interface(),
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
