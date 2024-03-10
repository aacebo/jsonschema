package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/formats"
	"os"
)

type namespace struct {
	schemas map[string]Schema
	formats map[string]Formatter
}

func New() *namespace {
	return &namespace{
		schemas: map[string]Schema{},
		formats: map[string]Formatter{
			"date-time": formats.DateTime,
			"email":     formats.Email,
			"ipv4":      formats.IPv4,
			"ipv6":      formats.IPv6,
			"uri":       formats.URI,
			"uuid":      formats.UUID,
		},
	}
}

func (self namespace) HasFormat(name string) bool {
	_, ok := self.formats[name]
	return ok
}

func (self *namespace) AddFormat(name string, format Formatter) *namespace {
	self.formats[name] = format
	return self
}

func (self namespace) Format(name string, input string) error {
	format, ok := self.formats[name]

	if !ok {
		return errors.New(fmt.Sprintf(`format "%s" does not exist`, name))
	}

	return format(input)
}

func (self namespace) HasSchema(id string) bool {
	_, ok := self.schemas[id]
	return ok
}

func (self namespace) GetSchema(id string) Schema {
	schema, ok := self.schemas[id]

	if !ok {
		return nil
	}

	return schema
}

func (self *namespace) AddSchema(schema Schema) *namespace {
	id := schema.GetID()

	if id == "" {
		panic(`"$id" is required for top level schemas`)
	}

	self.schemas[id] = schema
	return self
}

func (self *namespace) Read(path string) (Schema, error) {
	b, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	data := map[string]any{}
	err = json.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}

	schema, err := parse(data)

	if err != nil {
		return nil, err
	}

	id := schema.GetID()

	if id == "" {
		return nil, errors.New(`"$id" is required for top level schemas`)
	}

	self.schemas[id] = schema
	return schema, nil
}

// func (self *namespace) ReadDir(path string) ([]Schema, error) {
// 	entries, err := os.ReadDir(path)

// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, entry := range entries {
// 		os.
// 	}
// }

func (self namespace) Compile(id string) []SchemaCompileError {
	schema, ok := self.schemas[id]

	if !ok {
		panic(fmt.Sprintf(`schema "%s" not found`, id))
	}

	return schema.compile(self, "", "")
}

func (self namespace) Validate(id string, value any) []SchemaError {
	schema, ok := self.schemas[id]

	if !ok {
		panic(fmt.Sprintf(`schema "%s" not found`, id))
	}

	return schema.validate(self, "", "", value)
}
