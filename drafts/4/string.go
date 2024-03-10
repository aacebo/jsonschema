package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

// https://json-schema.org/understanding-json-schema/reference/string#format
type Formatter func(input string) error

// https://json-schema.org/understanding-json-schema/reference/string
type StringSchema struct {
	ID          *string    `json:"$id,omitempty"`         // https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
	Type        SchemaType `json:"type"`                  // https://json-schema.org/understanding-json-schema/reference/type
	Title       *string    `json:"title,omitempty"`       // https://json-schema.org/understanding-json-schema/reference/annotations
	Description *string    `json:"description,omitempty"` // https://json-schema.org/understanding-json-schema/reference/annotations
	Pattern     *string    `json:"pattern,omitempty"`     // https://json-schema.org/understanding-json-schema/reference/string#regexp
	Format      *string    `json:"format,omitempty"`      // https://json-schema.org/understanding-json-schema/reference/string#format
	MinLength   *int       `json:"minLength,omitempty"`   // https://json-schema.org/understanding-json-schema/reference/string#length
	MaxLength   *int       `json:"maxLength,omitempty"`   // https://json-schema.org/understanding-json-schema/reference/string#length
}

func (self StringSchema) GetID() string {
	if self.ID != nil {
		return *self.ID
	}

	return ""
}

func (self StringSchema) GetType() SchemaType {
	return self.Type
}

func (self StringSchema) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self StringSchema) compile(ns namespace, path string, key string) []SchemaCompileError {
	errors := []SchemaCompileError{}

	if key != "" {
		path = fmt.Sprintf("%s/%s", path, key)
	}

	if self.Type != SCHEMA_TYPE_STRING {
		errors = append(errors, SchemaCompileError{
			Path:    path,
			Keyword: "type",
			Message: fmt.Sprintf(`"type" must be "%s"`, SCHEMA_TYPE_STRING),
		})
	}

	if self.Pattern != nil {
		_, err := regexp.Compile(*self.Pattern)

		if err != nil {
			errors = append(errors, SchemaCompileError{
				Path:    path,
				Keyword: "pattern",
				Message: fmt.Sprintf(
					`"%s" is not a valid regex pattern`,
					*self.Pattern,
				),
			})
		}
	}

	if self.Format != nil {
		if !ns.HasFormat(*self.Format) {
			errors = append(errors, SchemaCompileError{
				Path:    path,
				Keyword: "format",
				Message: fmt.Sprintf(
					`format "%s" not found`,
					*self.Format,
				),
			})
		}
	}

	if self.MinLength != nil && *self.MinLength < 0 {
		errors = append(errors, SchemaCompileError{
			Path:    path,
			Keyword: "minLength",
			Message: `"minLength" must be non-negative`,
		})
	}

	if self.MaxLength != nil && *self.MaxLength < 0 {
		errors = append(errors, SchemaCompileError{
			Path:    path,
			Keyword: "maxLength",
			Message: `"maxLength" must be non-negative`,
		})
	}

	if self.MinLength != nil && self.MaxLength != nil && *self.MaxLength < *self.MinLength {
		errors = append(errors, SchemaCompileError{
			Path:    path,
			Keyword: "maxLength",
			Message: `"maxLength" must be greater than or equal to "minLength"`,
		})
	}

	return errors
}

func (self StringSchema) validate(ns namespace, path string, key string, value any) []SchemaError {
	errors := []SchemaError{}
	v := reflect.ValueOf(value)

	if key != "" {
		path = fmt.Sprintf("%s/%s", path, key)
	}

	if v.Kind() != reflect.String {
		errors = append(errors, SchemaError{
			Path:    path,
			Keyword: "type",
			Message: `should be type "string"`,
		})

		return errors
	}

	if self.Pattern != nil {
		exists := regexp.MustCompile(*self.Pattern).MatchString(v.String())

		if !exists {
			errors = append(errors, SchemaError{
				Path:    path,
				Keyword: "pattern",
				Message: fmt.Sprintf(
					`"%s" does not match pattern "%s"`,
					v.String(),
					*self.Pattern,
				),
			})
		}
	}

	if self.Format != nil {
		if err := ns.Format(*self.Format, v.String()); err != nil {
			errors = append(errors, SchemaError{
				Path:    path,
				Keyword: "format",
				Message: err.Error(),
			})
		}
	}

	if self.MinLength != nil {
		if *self.MinLength > len(v.String()) {
			errors = append(errors, SchemaError{
				Path:    path,
				Keyword: "minLength",
				Message: fmt.Sprintf(
					`length of "%d" is less than min length "%d"`,
					len(v.String()),
					*self.MinLength,
				),
			})
		}
	}

	if self.MaxLength != nil {
		if *self.MaxLength < len(v.String()) {
			errors = append(errors, SchemaError{
				Path:    path,
				Keyword: "maxLength",
				Message: fmt.Sprintf(
					`length of "%d" is greater than max length "%d"`,
					len(v.String()),
					*self.MaxLength,
				),
			})
		}
	}

	return errors
}

func parseString(data map[string]any) (StringSchema, error) {
	self := StringSchema{Type: SCHEMA_TYPE_STRING}

	if data == nil {
		return self, errors.New(`cannot parse "null" to "StringSchema"`)
	}

	if id, ok := data["$id"].(string); ok {
		self.ID = &id
	}

	if title, ok := data["title"].(string); ok {
		self.Title = &title
	}

	if description, ok := data["description"].(string); ok {
		self.Description = &description
	}

	if pattern, ok := data["pattern"].(string); ok {
		self.Pattern = &pattern
	}

	if minLength, ok := data["minLength"].(float64); ok {
		v := int(minLength)
		self.MinLength = &v
	}

	if maxLength, ok := data["maxLength"].(float64); ok {
		v := int(maxLength)
		self.MaxLength = &v
	}

	return self, nil
}
