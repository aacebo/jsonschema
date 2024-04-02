package jsonschema

import (
	"reflect"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/array#contains
func contains(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}
			config = coerce.Map(config)

			if config.Kind() != reflect.Map {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `should be a "Schema"`,
				})

				return errs
			}

			return ns.compile(
				ctx.ID,
				ctx.Path,
				config.Interface().(map[string]any),
			)
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || (value.Kind() != reflect.Slice && value.Kind() != reflect.Array) {
				return errs
			}

			config = coerce.Map(config)

			for i := 0; i < value.Len(); i++ {
				index := value.Index(i).Elem()
				_errs := ns.validate(
					ctx.ID,
					ctx.Path,
					config.Interface().(map[string]any),
					index.Interface(),
				)

				if len(_errs) == 0 {
					return errs
				}
			}

			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: key,
				Message: "should match one or more items",
			})

			return errs
		},
	}
}
