package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/object#propertyNames
func propertyNames(key string) Keyword {
	return Keyword{
		Default: map[string]any{},
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.Map {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "Schema"`,
				})

				return errs
			}

			_errs := ns.compile(
				ctx.ID,
				fmt.Sprintf("%s/%s", ctx.Path, key),
				config.Interface().(map[string]any),
			)

			if len(_errs) > 0 {
				errs = append(errs, _errs...)
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if value.Kind() != reflect.Map {
				return errs
			}

			iter := value.MapRange()

			for iter.Next() {
				_key := iter.Key()
				_errs := ns.validate(
					ctx.ID,
					fmt.Sprintf("%s/%s", ctx.Path, _key.String()),
					config.Interface().(map[string]any),
					_key.Interface(),
				)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			return errs
		},
	}
}
