package jsonschema

import "reflect"

// https://json-schema.org/understanding-json-schema/reference/combining#not
func not(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))

			if config.Kind() != reflect.Map {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `should be a "Schema"`,
				})

				return errs
			}

			return ns.compile(ctx.Path, config.Interface().(map[string]any))
		},
		Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
			errs := []SchemaError{}
			config := reflect.Indirect(reflect.ValueOf(ctx.Value))
			_errs := ns.validate(ctx.Path, config.Interface().(map[string]any), input)

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
