package jsonschema

// https://json-schema.org/understanding-json-schema/reference/string#format
type Formatter func(input string) error
