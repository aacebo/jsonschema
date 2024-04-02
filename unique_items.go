package jsonschema

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array#uniqueItems
func uniqueItems(key string) Keyword {
	return Keyword{
		Default: false,
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.Bool {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `should be a "boolean"`,
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || (value.Kind() != reflect.Slice && value.Kind() != reflect.Array) {
				return errs
			}

			if config.Bool() {
				items := map[string]int{}

				for i := 0; i < value.Len(); i++ {
					index := value.Index(i)
					b, err := json.Marshal(index.Interface())

					if err != nil {
						errs = append(errs, SchemaError{
							Path:    ctx.Path,
							Keyword: key,
							Message: err.Error(),
						})

						continue
					}

					json := string(b)
					items[json]++

					if items[json] > 1 {
						errs = append(errs, SchemaError{
							Path:    fmt.Sprintf("%s/%d", ctx.Path, i),
							Keyword: key,
							Message: "duplicate item",
						})
					}
				}
			}

			return errs
		},
	}
}
