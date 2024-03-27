package jsonschema

import (
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
		return reflect.Array
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
var schemaType = Keyword{
	compile: func(ns *Namespace, ctx Context) []SchemaError {
		errs := []SchemaError{}

		switch v := ctx.Value.(type) {
		case string:
			if !SchemaType(v).Valid() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: "type",
					Message: `must be a valid "SchemaType"`,
				})
			}

			break
		case []string:
			for _, item := range v {
				if !SchemaType(item).Valid() {
					errs = append(errs, SchemaError{
						Path:    ctx.Path,
						Keyword: "type",
						Message: `must be a valid "SchemaType"`,
					})

					break
				}
			}

			break
		default:
			errs = append(errs, SchemaError{
				Path:    ctx.Path,
				Keyword: "type",
				Message: `must be a "string" or "[]string"`,
			})
		}

		return errs
	},
	validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}

		return errs
	},
}
