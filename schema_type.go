package jsonschema

import (
	"fmt"
	"reflect"
)

type SchemaType string

const (
	SCHEMA_TYPE_STRING  SchemaType = "string"
	SCHEMA_TYPE_NUMBER  SchemaType = "number"
	SCHEMA_TYPE_NULL    SchemaType = "null"
	SCHEMA_TYPE_INTEGER SchemaType = "integer"
	SCHEMA_TYPE_BOOLEAN SchemaType = "boolean"
	SCHEMA_TYPE_ARRAY   SchemaType = "array"
	SCHEMA_TYPE_OBJECT  SchemaType = "object"
)

func (self SchemaType) Valid() bool {
	switch self {
	case SCHEMA_TYPE_ARRAY:
		fallthrough
	case SCHEMA_TYPE_BOOLEAN:
		fallthrough
	case SCHEMA_TYPE_INTEGER:
		fallthrough
	case SCHEMA_TYPE_NULL:
		fallthrough
	case SCHEMA_TYPE_NUMBER:
		fallthrough
	case SCHEMA_TYPE_OBJECT:
		fallthrough
	case SCHEMA_TYPE_STRING:
		return true
	}

	return false
}

func (self SchemaType) Kind() reflect.Kind {
	switch self {
	case SCHEMA_TYPE_ARRAY:
		return reflect.Slice
	case SCHEMA_TYPE_BOOLEAN:
		return reflect.Bool
	case SCHEMA_TYPE_INTEGER:
		return reflect.Int
	case SCHEMA_TYPE_NUMBER:
		return reflect.Float64
	case SCHEMA_TYPE_OBJECT:
		return reflect.Map
	case SCHEMA_TYPE_STRING:
		return reflect.String
	}

	return reflect.Invalid
}

// https://json-schema.org/understanding-json-schema/reference/type
func schemaType(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			switch config.Kind() {
			case reflect.String:
				if !SchemaType(config.String()).Valid() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: key,
						Message: `must be a valid "SchemaType"`,
					})
				}

				break
			case reflect.Slice:
				for i := 0; i < config.Len(); i++ {
					index := config.Index(i).Elem()

					if index.Kind() != reflect.String || !SchemaType(index.String()).Valid() {
						errs = append(errs, SchemaError{
							Path:    fmt.Sprintf("%s/%s/%d", ctx.Path, key, i),
							Keyword: key,
							Message: `must be a valid "SchemaType"`,
						})

						break
					}
				}

				break
			default:
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string" or "[]string"`,
				})
			}

			return errs
		},
		Validate: func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError {
			errs := []SchemaError{}
			types := []string{}

			if !value.IsValid() {
				return errs
			}

			if config.Kind() == reflect.String {
				types = append(types, config.String())
			} else {
				for i := 0; i < config.Len(); i++ {
					index := config.Index(i).Elem()
					types = append(types, index.String())
				}
			}

			for _, t := range types {
				if SchemaType(t).Kind() == value.Kind() {
					return errs
				}
			}

			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: key,
				Message: fmt.Sprintf(
					`"%s" should be one of %v`,
					value.Kind().String(),
					types,
				),
			})

			return errs
		},
	}
}
