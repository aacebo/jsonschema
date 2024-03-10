package jsonschema

import (
	"errors"
	"fmt"
	"jsonschema/core"
)

type Schema interface {
	core.Schema

	compile(ns namespace, path string, key string) []core.SchemaError
	validate(ns namespace, path string, key string, value any) []core.SchemaError
}

func parse(data map[string]any) (Schema, error) {
	if data == nil {
		return nil, nil
	}

	t, ok := data["type"].(string)

	if !ok {
		return nil, errors.New("schema type is required and must be a string")
	}

	switch core.SchemaType(t) {
	case core.SCHEMA_TYPE_STRING:
		return parseString(data)
	case core.SCHEMA_TYPE_NUMBER:
		return parseNumber(data)
	case core.SCHEMA_TYPE_ARRAY:
		return parseArray(data)
	case core.SCHEMA_TYPE_NULL:
		return parseNull(data)
	}

	return nil, errors.New(fmt.Sprintf(`invalid schema type "%s"`, t))
}
