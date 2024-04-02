package jsonschema

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/aacebo/jsonschema/coerce"
)

// https://json-schema.org/understanding-json-schema/reference/object#patternProperties
func patternProperties(key string) Keyword {
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
				_, err := regexp.Compile(_key.String())

				if value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
					value = value.Elem()
				}

				value = coerce.Map(value)

				if err != nil {
					errs = append(errs, SchemaError{
						Path:    path,
						Keyword: key,
						Message: "must be a valid regular expression",
					})

					continue
				}

				if value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
					value = value.Elem()
				}

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
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if !value.IsValid() || (value.Kind() != reflect.Map && value.Kind() != reflect.Struct) {
				return errs
			}

			iter := config.MapRange()

			for iter.Next() {
				_key := iter.Key()
				_schema := reflect.Indirect(iter.Value())
				expr := regexp.MustCompile(_key.String())
				path := fmt.Sprintf("%s/%s", ctx.Path, _key.String())

				if _schema.Kind() == reflect.Pointer || _schema.Kind() == reflect.Interface {
					_schema = _schema.Elem()
				}

				_schema = coerce.Map(_schema)

				if _schema.Kind() != reflect.Map {
					errs = append(errs, SchemaError{
						Path:    path,
						Keyword: "type",
						Message: `should be a "Schema"`,
					})

					continue
				}

				if value.Kind() == reflect.Map {
					_iter := value.MapRange()

					for _iter.Next() {
						if expr.MatchString(_iter.Key().String()) {
							_errs := ns.validate(
								ctx.ID,
								path,
								_schema.Interface().(map[string]any),
								_iter.Value().Interface(),
							)

							if len(_errs) > 0 {
								errs = append(errs, _errs...)
							}
						}
					}
				}

				if value.Kind() == reflect.Struct {
					for i := 0; i < value.NumField(); i++ {
						field := value.Field(i)
						name := coerce.StructFieldName(value.Type().Field(i))

						if expr.MatchString(name) {
							_errs := ns.validate(
								ctx.ID,
								path,
								_schema.Interface().(map[string]any),
								field.Interface(),
							)

							if len(_errs) > 0 {
								errs = append(errs, _errs...)
							}
						}
					}
				}
			}

			return errs
		},
	}
}
