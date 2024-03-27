package jsonschema

import (
	"errors"
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

func (self SchemaType) Validate(value any) error {
	switch self {
	case SCHEMA_TYPE_ARRAY:
		if _, ok := value.([]any); !ok {
			return errors.New(`must be an "array"`)
		}

		break
	case SCHEMA_TYPE_BOOLEAN:
		if _, ok := value.(bool); !ok {
			return errors.New(`must be a "bool"`)
		}

		break
	case SCHEMA_TYPE_INTEGER:
		if _, ok := value.(int); !ok {
			return errors.New(`must be a "int"`)
		}

		break
	case SCHEMA_TYPE_NULL:
		if value != nil {
			return errors.New(`must be "null"`)
		}

		break
	case SCHEMA_TYPE_NUMBER:
		if _, ok := value.(float64); !ok {
			return errors.New(`must be a "float"`)
		}

		break
	case SCHEMA_TYPE_OBJECT:
		if _, ok := value.(map[string]any); !ok {
			return errors.New(`must be a "map"`)
		}

		break
	case SCHEMA_TYPE_STRING:
		if _, ok := value.(string); !ok {
			return errors.New(`must be a "string"`)
		}
	}

	return nil
}

// https://json-schema.org/understanding-json-schema/reference/type
var schemaType = Keyword{
	Compile: func(ns *Namespace, ctx Context) []SchemaError {
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
		case []any:
			for i, item := range v {
				str, ok := item.(string)

				if !ok {
					errs = append(errs, SchemaError{
						Path:    fmt.Sprintf("%s/type/%d", ctx.Path, i),
						Keyword: "type",
						Message: `must be a string`,
					})

					continue
				}

				if !SchemaType(str).Valid() {
					errs = append(errs, SchemaError{
						Path:    fmt.Sprintf("%s/type/%d", ctx.Path, i),
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
	Validate: func(ns *Namespace, ctx Context, input any) []SchemaError {
		errs := []SchemaError{}
		types := []string{}
		t, ok := ctx.Value.(string)
		value := reflect.Indirect(reflect.ValueOf(input))

		if !ok {
			ts, _ := ctx.Value.([]any)

			for _, t := range ts {
				types = append(types, t.(string))
			}
		} else {
			types = []string{t}
		}

		for _, t := range types {
			if SchemaType(t).Validate(input) == nil {
				return errs
			}
		}

		errs = append(errs, SchemaError{
			Path:    ctx.Path,
			Keyword: "type",
			Message: fmt.Sprintf(
				`"%s" should be one of %v`,
				value.Kind().String(),
				types,
			),
		})

		return errs
	},
}
