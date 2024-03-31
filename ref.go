package jsonschema

import (
	"reflect"
)

// https://json-schema.org/understanding-json-schema/structuring#dollarref
func ref(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.String {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})

				return errs
			}

			schema, err := ns.Resolve(ctx.ID, config.String())

			if err != nil {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: err.Error(),
				})
			}

			_errs := ns.compile(schema.ID(), ctx.Path, schema)

			if len(_errs) > 0 {
				errs = append(errs, _errs...)
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}
			schema, err := ns.Resolve(ctx.ID, config.String())

			if err != nil {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: err.Error(),
				})
			}

			_errs := ns.validate(schema.ID(), ctx.Path, schema, value.Interface())

			if len(_errs) > 0 {
				errs = append(errs, _errs...)
			}

			return errs
		},
	}
}
