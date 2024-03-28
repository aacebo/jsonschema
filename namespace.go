package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/formats"
	"os"
)

type Namespace struct {
	schemas  map[string]Schema
	keywords map[string]Keyword
	formats  map[string]Formatter
}

func New() *Namespace {
	return &Namespace{
		schemas: map[string]Schema{},
		keywords: map[string]Keyword{
			"$id":              id,
			"id":               id,
			"$schema":          schemaSpec,
			"type":             schemaType,
			"title":            title,
			"description":      description,
			"dependencies":     dependencies,
			"pattern":          pattern,
			"format":           format,
			"minLength":        minLength,
			"maxLength":        maxLength,
			"multipleOf":       multipleOf,
			"minimum":          minimum,
			"maximum":          maximum,
			"exclusiveMinimum": exclusiveMinimum,
			"exclusiveMaximum": exclusiveMaximum,
			"enum":             enum,
			"items":            items,
			"additionalItems":  additionalItems,
			"anyOf":            anyOf,
			"allOf":            allOf,
			"oneOf":            oneOf,
			"not":              not,
			"default":          _default,
			"const":            _const,
			"$comment":         comment,
		},
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

func (self Namespace) HasFormat(name string) bool {
	_, ok := self.formats[name]
	return ok
}

func (self *Namespace) AddFormat(name string, format Formatter) *Namespace {
	self.formats[name] = format
	return self
}

func (self Namespace) Format(name string, input string) error {
	format, ok := self.formats[name]

	if !ok {
		return errors.New(fmt.Sprintf(`format "%s" does not exist`, name))
	}

	return format(input)
}

func (self Namespace) HasSchema(id string) bool {
	_, ok := self.schemas[id]
	return ok
}

func (self Namespace) GetSchema(id string) Schema {
	schema, ok := self.schemas[id]

	if !ok {
		return nil
	}

	return schema
}

func (self *Namespace) AddSchema(schema Schema) *Namespace {
	self.schemas[schema.ID()] = schema
	return self
}

func (self *Namespace) Keyword(name string, keyword Keyword) *Namespace {
	self.keywords[name] = keyword
	return self
}

func (self *Namespace) Compile(id string) []SchemaError {
	schema, ok := self.schemas[id]

	if !ok {
		return []SchemaError{}
	}

	return self.compile("", schema)
}

func (self *Namespace) Validate(id string, value any) []SchemaError {
	schema, ok := self.schemas[id]

	if !ok {
		return []SchemaError{}
	}

	return self.validate("", schema, value)
}

func (self *Namespace) Read(path string) (Schema, error) {
	b, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	schema := Schema{}
	err = json.Unmarshal(b, &schema)

	if err != nil {
		return nil, err
	}

	id := schema.ID()
	self.schemas[id] = schema
	return schema, nil
}

func (self *Namespace) compile(path string, schema Schema) []SchemaError {
	errs := []SchemaError{}

	for key, value := range schema {
		keyword, ok := self.keywords[key]

		if !ok || value == nil || keyword.Compile == nil {
			continue
		}

		err := keyword.Compile(self, Context{
			Path:   path,
			Schema: schema,
			Value:  value,
		})

		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	return errs
}

func (self *Namespace) validate(path string, schema Schema, value any) []SchemaError {
	errs := []SchemaError{}
	defaultValue, _ := schema["default"]

	if value == nil {
		value = defaultValue
	}

	for key, keyword := range self.keywords {
		config, ok := schema[key]

		if !ok && keyword.Default != nil {
			config = keyword.Default
		}

		if config == nil || keyword.Validate == nil {
			continue
		}

		err := keyword.Validate(self, Context{
			Path:   path,
			Schema: schema,
			Value:  config,
		}, value)

		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	return errs
}
