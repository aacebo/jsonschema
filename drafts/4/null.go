package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/core"
)

// https://json-schema.org/understanding-json-schema/reference/null
type NullSchema struct {
	ID          *string         `json:"$id,omitempty"`         // https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
	Type        core.SchemaType `json:"type"`                  // https://json-schema.org/understanding-json-schema/reference/type
	Title       *string         `json:"title,omitempty"`       // https://json-schema.org/understanding-json-schema/reference/annotations
	Description *string         `json:"description,omitempty"` // https://json-schema.org/understanding-json-schema/reference/annotations
}

func (self NullSchema) GetID() string {
	if self.ID != nil {
		return *self.ID
	}

	return ""
}

func (self NullSchema) GetType() core.SchemaType {
	return self.Type
}

func (self NullSchema) GetTitle() string {
	if self.Title != nil {
		return *self.Title
	}

	return ""
}

func (self NullSchema) GetDescription() string {
	if self.Description != nil {
		return *self.Description
	}

	return ""
}

func (self NullSchema) Value() any {
	return map[string]any{"type": "null"}
}

func (self NullSchema) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self NullSchema) compile(ns core.Namespace[Schema], id string, path string) []core.SchemaError {
	errors := []core.SchemaError{}

	if self.Type != core.SCHEMA_TYPE_NULL {
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "type",
			Message: fmt.Sprintf(`"type" must be "%s"`, core.SCHEMA_TYPE_NULL),
		})
	}

	return errors
}

func (self NullSchema) validate(ns core.Namespace[Schema], id string, path string, value any) []core.SchemaError {
	errors := []core.SchemaError{}

	if value != nil {
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "type",
			Message: fmt.Sprintf(`"type" must be "%s"`, core.SCHEMA_TYPE_NULL),
		})
	}

	return errors
}

func parseNull(data map[string]any) (NullSchema, error) {
	self := NullSchema{Type: core.SCHEMA_TYPE_NULL}

	if data == nil {
		return self, errors.New(`cannot parse "null" to "NullSchema"`)
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

	return self, nil
}
