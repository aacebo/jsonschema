package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/aacebo/jsonschema/formats"
)

type Namespace struct {
	schemas  map[string]Schema
	keywords map[string]func(string) Keyword
	formats  map[string]Formatter
}

func New() *Namespace {
	return &Namespace{
		schemas: map[string]Schema{},
		keywords: map[string]func(string) Keyword{
			"$id":                  id,
			"id":                   id,
			"$schema":              dialect,
			"$defs":                definitions,
			"definitions":          definitions,
			"$comment":             comment,
			"$ref":                 ref,
			"type":                 schemaType,
			"title":                title,
			"description":          description,
			"anyOf":                anyOf,
			"allOf":                allOf,
			"oneOf":                oneOf,
			"not":                  not,
			"default":              _default,
			"const":                _const,
			"dependencies":         dependencies,
			"pattern":              pattern,
			"format":               format,
			"minLength":            minLength,
			"maxLength":            maxLength,
			"multipleOf":           multipleOf,
			"minimum":              minimum,
			"maximum":              maximum,
			"exclusiveMinimum":     exclusiveMinimum,
			"exclusiveMaximum":     exclusiveMaximum,
			"enum":                 enum,
			"items":                items,
			"minItems":             minItems,
			"maxItems":             maxItems,
			"additionalItems":      additionalItems,
			"uniqueItems":          uniqueItems,
			"contains":             contains,
			"properties":           properties,
			"patternProperties":    patternProperties,
			"additionalProperties": additionalProperties,
			"propertyNames":        propertyNames,
			"minProperties":        minProperties,
			"maxProperties":        maxProperties,
			"required":             required,
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

func (self *Namespace) AddKeyword(name string, keyword Keyword) *Namespace {
	self.keywords[name] = func(_ string) Keyword {
		return keyword
	}

	return self
}

func (self *Namespace) Compile(schema Schema) []SchemaError {
	id := schema.ID()
	_, ok := self.schemas[id]

	if !ok {
		self.schemas[id] = schema
	}

	return self.compile(id, "", schema)
}

func (self *Namespace) Validate(schema Schema, value any) []SchemaError {
	id := schema.ID()
	_, ok := self.schemas[id]

	if !ok {
		self.schemas[id] = schema
	}

	return self.validate(id, "", schema, value)
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
		factory, ok := self.keywords[key]

		if !ok {
			errs = append(errs, SchemaError{
				Path:    path,
				Keyword: key,
				Message: "invalid keyword",
			})

			continue
		}

		if config == nil {
			continue
		}

		keyword := factory(key)

		if keyword.Compile == nil {
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

	for key, factory := range self.keywords {
		keyword := factory(key)
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
