package jsonschema

import (
	"errors"
	"fmt"
	"jsonschema/core"
)

type Schema interface {
	core.Schema
	compile(ns core.Namespace[Schema], id string, path string) []core.SchemaError
	validate(ns core.Namespace[Schema], id string, path string, value any) []core.SchemaError
}

func parse(data map[string]any) (Schema, error) {
	if data == nil {
		return nil, nil
	}

	ref, ok := data["$ref"].(string)

	if ok {
		return RefSchema{ref}, nil
	}

	t, ok := data["type"]

	if !ok {
		return nil, errors.New("schema type is required and must be a string")
	}

	switch core.SchemaType(fmt.Sprint(t)) {
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
