package jsonschema

import "reflect"

// https://json-schema.org/understanding-json-schema/reference/combining#not
func not(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

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
			_errs := ns.validate(
				ctx.ID,
				ctx.Path,
				config.Interface().(map[string]any),
				value.Interface(),
			)

			if len(_errs) == 0 {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: "should not match schema",
				})
			}

			return errs
		},
	}
}
