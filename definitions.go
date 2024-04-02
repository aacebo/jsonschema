package jsonschema

import (
	"fmt"
	"reflect"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/structuring#defs
func definitions(key string) Keyword {
	return Keyword{
		Default: Schema{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.Map {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "map"`,
				})

				return errs
			}

			iter := config.MapRange()

			for iter.Next() {
				_key := iter.Key()
				value := reflect.Indirect(iter.Value())
				path := fmt.Sprintf("%s/%s/%s", ctx.Path, key, _key.String())

				if value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
					value = value.Elem()
				}

				value = coerce.Map(value)

				if value.Kind() != reflect.Map {
					errs = append(errs, SchemaError{
						Path:    path,
						Keyword: "type",
						Message: `should be a "Schema"`,
					})

					continue
				}

				_errs := ns.compile(
					ctx.ID,
					path,
					value.Interface().(map[string]any),
				)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			return errs
		},
	}
}
