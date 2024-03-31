package jsonschema

import (
	"fmt"
	"reflect"
	"regexp"
)

// https://json-schema.org/understanding-json-schema/reference/object#additionalproperties
func additionalProperties(key string) Keyword {
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

			if !value.IsValid() || value.Kind() != reflect.Map {
				return errs
			}

			properties := reflect.Indirect(reflect.ValueOf(ctx.Schema["properties"]))
			patternProperties := reflect.Indirect(reflect.ValueOf(ctx.Schema["patternProperties"]))
			iter := value.MapRange()

			for iter.Next() {
				_key := iter.Key()
				_value := iter.Value()

				if properties.IsValid() {
					if properties.MapIndex(_key).IsValid() && !properties.MapIndex(_key).IsZero() {
						continue
					}
				}

				if patternProperties.IsValid() {
					_iter := patternProperties.MapRange()
					_match := false

					for _iter.Next() {
						expr := regexp.MustCompile(_iter.Key().String())
						_match = expr.MatchString(_key.String())

						if _match {
							break
						}
					}

					if _match {
						continue
					}
				}

				fmt.Println(_key.String())

				switch config.Kind() {
				case reflect.Bool:
					if !config.Bool() {
						errs = append(errs, SchemaError{
							Path:    fmt.Sprintf("%s/%s", ctx.Path, _key.String()),
							Keyword: key,
							Message: "too many properties",
						})
					}

					break
				case reflect.Map:
					_errs := ns.validate(
						ctx.ID,
						fmt.Sprintf("%s/%s", ctx.Path, _key.String()),
						config.Interface().(map[string]any),
						_value.Interface(),
					)

					if len(_errs) > 0 {
						errs = append(errs, _errs...)
					}

					break
				}
			}

			return errs
		},
	}
}
