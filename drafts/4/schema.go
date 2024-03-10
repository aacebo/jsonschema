package jsonschema

import (
	"errors"
	"fmt"
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

type Schema interface {
	GetID() string
	GetType() SchemaType

	compile(ns namespace, path string, key string) []SchemaCompileError
	validate(ns namespace, path string, key string, value any) []SchemaError
}

func parse(data map[string]any) (Schema, error) {
	if data == nil {
		return nil, nil
	}

	t, ok := data["type"].(string)

	if !ok {
		return nil, errors.New("schema type is required and must be a string")
	}

	switch SchemaType(t) {
	case SCHEMA_TYPE_STRING:
		return parseString(data)
	case SCHEMA_TYPE_NUMBER:
		return parseNumber(data)
	case SCHEMA_TYPE_ARRAY:
		return parseArray(data)
	case SCHEMA_TYPE_NULL:
		return parseNull(data)
	}

	return nil, errors.New(fmt.Sprintf(`invalid schema type "%s"`, t))
}
