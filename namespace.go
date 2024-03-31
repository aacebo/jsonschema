package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/formats"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
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
			"$id":                  id("$id"),
			"id":                   id("id"),
			"$schema":              schemaSpec("$schema"),
			"$defs":                definitions("$defs"),
			"definitions":          definitions("definitions"),
			"type":                 schemaType("type"),
			"title":                title("title"),
			"description":          description("description"),
			"dependencies":         dependencies("dependencies"),
			"pattern":              pattern("pattern"),
			"format":               format("format"),
			"minLength":            minLength("minLength"),
			"maxLength":            maxLength("maxLength"),
			"multipleOf":           multipleOf("multipleOf"),
			"minimum":              minimum("minimum"),
			"maximum":              maximum("maximum"),
			"exclusiveMinimum":     exclusiveMinimum("exclusiveMinimum"),
			"exclusiveMaximum":     exclusiveMaximum("exclusiveMaximum"),
			"enum":                 enum("enum"),
			"items":                items("items"),
			"minItems":             minItems("minItems"),
			"maxItems":             maxItems("maxItems"),
			"additionalItems":      additionalItems("additionalItems"),
			"uniqueItems":          uniqueItems("uniqueItems"),
			"contains":             contains("contains"),
			"properties":           properties("properties"),
			"patternProperties":    patternProperties("patternProperties"),
			"additionalProperties": additionalProperties("additionalProperties"),
			"anyOf":                anyOf("anyOf"),
			"allOf":                allOf("allOf"),
			"oneOf":                oneOf("oneOf"),
			"not":                  not("not"),
			"default":              _default("default"),
			"const":                _const("const"),
			"$comment":             comment("$comment"),
			"$ref":                 ref("$ref"),
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

	return self.compile(schema.ID(), "", schema)
}

func (self *Namespace) Validate(id string, value any) []SchemaError {
	schema, ok := self.schemas[id]

	if !ok {
		return []SchemaError{}
	}

	return self.validate(schema.ID(), "", schema, value)
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

func (self *Namespace) Resolve(id string, path string) (Schema, error) {
	url, path, err := self.parseRef(id, path)

	if err != nil {
		return nil, err
	}

	schema := self.GetSchema(url)

	if schema == nil {
		res, err := http.Get(url)

		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(res.Body).Decode(&schema)

		if err != nil {
			return nil, errors.New("failed to parse remote schema to json")
		}

		self.schemas[url] = schema
	}

	if path != "" {
		var curr any = schema
		parts := strings.Split(path, "/")

		for _, part := range parts {
			if part == "" || part == "#" {
				continue
			}

			switch v := curr.(type) {
			case Schema:
				curr = v[part]
				break
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

		schema = curr.(map[string]any)
	}

	return schema, nil
}

func (self *Namespace) compile(id string, path string, schema Schema) []SchemaError {
	errs := []SchemaError{}

	for key, config := range schema {
		keyword, ok := self.keywords[key]

		if !ok || config == nil || keyword.Compile == nil {
			continue
		}

		err := keyword.Compile(
			self,
			Context{id, path, schema},
			reflect.Indirect(reflect.ValueOf(config)),
		)

		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	return errs
}

func (self *Namespace) validate(id string, path string, schema Schema, value any) []SchemaError {
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

		err := keyword.Validate(
			self,
			Context{id, path, schema},
			reflect.Indirect(reflect.ValueOf(config)),
			reflect.Indirect(reflect.ValueOf(value)),
		)

		if len(err) > 0 {
			errs = append(errs, err...)
		}
	}

	return errs
}

func (self *Namespace) parseRef(id string, path string) (string, string, error) {
	if path == "" {
		return id, "", nil
	}

	refURL, err := url.Parse(path)

	if err != nil {
		return "", "", err
	}

	if refURL.IsAbs() {
		return path, refURL.Fragment, nil
	}

	baseURL, err := url.Parse(id)

	if err != nil {
		return "", "", err
	}

	if strings.HasPrefix(refURL.String(), "#") {
		return baseURL.String(), refURL.String(), nil
	}

	return baseURL.ResolveReference(refURL).String(), "", nil
}
