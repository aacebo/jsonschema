package jsonschema

import (
	"fmt"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/object#properties
func properties(key string) Keyword {
	return Keyword{
		Default: map[string]any{},
		Compile: definitions(key).Compile,
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || value.Kind() != reflect.Map {
				return errs
			}

			iter := config.MapRange()

			for iter.Next() {
				_key := iter.Key()
				_schema := reflect.Indirect(iter.Value())
				_value := value.MapIndex(_key)
				path := fmt.Sprintf("%s/%s", ctx.Path, _key.String())

				if _value.IsZero() {
					continue
				}

				if _schema.Kind() == reflect.Pointer || _schema.Kind() == reflect.Interface {
					_schema = _schema.Elem()
				}

				if _schema.Kind() != reflect.Map {
					errs = append(errs, SchemaError{
						Path:    path,
						Keyword: "type",
						Message: `should be a "Schema"`,
					})

					continue
				}

				_errs := ns.validate(
					ctx.ID,
					path,
					_schema.Interface().(map[string]any),
					_value.Interface(),
				)

				if len(_errs) > 0 {
					errs = append(errs, _errs...)
				}
			}

			return errs
		},
	}
}
