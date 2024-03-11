package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/core"
	"jsonschema/formats"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type namespace struct {
	schemas map[string]Schema
	formats map[string]core.Formatter
}

func New() *namespace {
	return &namespace{
		schemas: map[string]Schema{},
		formats: map[string]core.Formatter{
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

func (self *namespace) AddFormat(name string, format core.Formatter) core.Namespace[Schema] {
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

func (self *namespace) AddSchema(schema Schema) core.Namespace[Schema] {
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

func (self *namespace) Resolve(url string, path string) (Schema, error) {
	schema := self.GetSchema(url)

	if schema == nil {
		res, err := http.Get(url)

		if err != nil {
			return nil, err
		}

		data := map[string]any{}
		err = json.NewDecoder(res.Body).Decode(&data)

		if err != nil {
			return nil, err
		}

		schema, err = parse(data)

		if err != nil {
			return nil, err
		}

		self.schemas[url] = schema
	}

	if path != "" {
		parts := strings.Split(path, "/")
		b, err := json.Marshal(schema)

		if err != nil {
			return nil, err
		}

		var curr any
		err = json.Unmarshal(b, &curr)

		if err != nil {
			return nil, err
		}

		for _, part := range parts {
			if part == "" || part == "#" {
				continue
			}

			switch v := curr.(type) {
			case map[string]any:
				curr = v[part]
				break
			case []any:
				i, err := strconv.Atoi(part)

				if err != nil {
					return nil, err
				}

				curr = v[i]
				break
			default:
				curr = nil
				break
			}

			if curr == nil {
				return nil, errors.New(fmt.Sprintf(
					`ref path "%s" not found`,
					path,
				))
			}
		}

		return parse(curr.(map[string]any))
	}

	return schema, nil
}

func (self *namespace) Compile(id string) []core.SchemaError {
	schema, ok := self.schemas[id]

	if !ok {
		panic(fmt.Sprintf(`schema "%s" not found`, id))
	}

	return schema.compile(self, "", "")
}

func (self *namespace) Validate(id string, value any) []core.SchemaError {
	schema, ok := self.schemas[id]

	if !ok {
		panic(fmt.Sprintf(`schema "%s" not found`, id))
	}

	return schema.validate(self, "", "", value)
}
