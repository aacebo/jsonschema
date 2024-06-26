package jsonschema

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/object#additionalproperties
func additionalProperties(key string) Keyword {
	return Keyword{
		Default: Schema{},
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

			if !value.IsValid() || (value.Kind() != reflect.Map && value.Kind() != reflect.Struct) {
				return errs
			}

			config = coerce.Map(config)
			properties := reflect.Indirect(reflect.ValueOf(ctx.Schema["properties"]))
			patternProperties := reflect.Indirect(reflect.ValueOf(ctx.Schema["patternProperties"]))

			switch value.Kind() {
			case reflect.Map:
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
			case reflect.Struct:
				for i := 0; i < value.NumField(); i++ {
					field := value.Field(i)
					name := coerce.StructFieldName(value.Type().Field(i))

					if properties.IsValid() {
						if properties.MapIndex(reflect.ValueOf(name)).IsValid() && !properties.MapIndex(reflect.ValueOf(name)).IsZero() {
							continue
						}
					}

					if patternProperties.IsValid() {
						_iter := patternProperties.MapRange()
						_match := false

						for _iter.Next() {
							expr := regexp.MustCompile(_iter.Key().String())
							_match = expr.MatchString(name)

							if _match {
								break
							}
						}

						if _match {
							continue
						}
					}

					switch config.Kind() {
					case reflect.Bool:
						if !config.Bool() {
							errs = append(errs, SchemaError{
								Path:    fmt.Sprintf("%s/%s", ctx.Path, name),
								Keyword: key,
								Message: "too many properties",
							})
						}

						break
					case reflect.Map:
						_errs := ns.validate(
							ctx.ID,
							fmt.Sprintf("%s/%s", ctx.Path, name),
							config.Interface().(map[string]any),
							field.Interface(),
						)

						if len(_errs) > 0 {
							errs = append(errs, _errs...)
						}

						break
					}
				}
			}

			return errs
		},
	}
}
