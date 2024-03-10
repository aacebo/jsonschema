package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// https://json-schema.org/understanding-json-schema/reference/array
type ArraySchema struct {
	ID              *string               `json:"$id,omitempty"`             // https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
	Type            SchemaType            `json:"type"`                      // https://json-schema.org/understanding-json-schema/reference/type
	Title           *string               `json:"title,omitempty"`           // https://json-schema.org/understanding-json-schema/reference/annotations
	Description     *string               `json:"description,omitempty"`     // https://json-schema.org/understanding-json-schema/reference/annotations
	Items           *ArrayItems           `json:"items,omitempty"`           // https://json-schema.org/understanding-json-schema/reference/array#items
	AdditionalItems *ArrayAdditionalItems `json:"additionalItems,omitempty"` // https://json-schema.org/understanding-json-schema/reference/array#additionalitems
	MinItems        *int                  `json:"minItems,omitempty"`        // https://json-schema.org/understanding-json-schema/reference/array#length
	MaxItems        *int                  `json:"maxItems,omitempty"`        // https://json-schema.org/understanding-json-schema/reference/array#length
	UniqueItems     *bool                 `json:"uniqueItems,omitempty"`     // https://json-schema.org/understanding-json-schema/reference/array#uniqueItems
}

func (self ArraySchema) GetID() string {
	if self.ID != nil {
		return *self.ID
	}

	return ""
}

func (self ArraySchema) GetType() SchemaType {
	return self.Type
}

func (self ArraySchema) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self ArraySchema) compile(ns namespace, path string, key string) []SchemaCompileError {
	errors := []SchemaCompileError{}

	if key != "" {
		path = fmt.Sprintf("%s/%s", path, key)
	}

	if self.Type != SCHEMA_TYPE_ARRAY {
		errors = append(errors, SchemaCompileError{
			Path:    path,
			Keyword: "type",
			Message: fmt.Sprintf(`"type" must be "%s"`, SCHEMA_TYPE_ARRAY),
		})
	}

	if self.Items != nil {
		if self.Items.One != nil {
			if errs := self.Items.One.compile(ns, path, "items"); len(errs) > 0 {
				errors = append(errors, errs...)
			}
		}

		if self.Items.Many != nil {
			for i, schema := range self.Items.Many {
				if errs := schema.compile(ns, path, fmt.Sprintf("items/%d", i)); len(errs) > 0 {
					errors = append(errors, errs...)
				}
			}
		}
	}

	if self.AdditionalItems != nil {
		if self.AdditionalItems.Schema != nil {
			errs := self.AdditionalItems.Schema.compile(ns, path, "additionalItems")

			if len(errs) > 0 {
				errors = append(errors, errs...)
			}
		}
	}

	if self.MinItems != nil {
		if *self.MinItems < 0 {
			errors = append(errors, SchemaCompileError{
				Path:    path,
				Keyword: "minItems",
				Message: "must be non-negative",
			})
		}
	}

	if self.MaxItems != nil {
		if *self.MaxItems < 0 {
			errors = append(errors, SchemaCompileError{
				Path:    path,
				Keyword: "maxItems",
				Message: "must be non-negative",
			})
		}
	}

	if self.MinItems != nil && self.MaxItems != nil && *self.MaxItems < *self.MinItems {
		errors = append(errors, SchemaCompileError{
			Path:    path,
			Keyword: "maxItems",
			Message: `must be greater than or equal to "minItems"`,
		})
	}

	return errors
}

func (self ArraySchema) validate(ns namespace, path string, key string, value any) []SchemaError {
	errors := []SchemaError{}
	v := reflect.ValueOf(value)

	if key != "" {
		path = fmt.Sprintf("%s/%s", path, key)
	}

	if v.Kind() != reflect.Slice {
		errors = append(errors, SchemaError{
			Path:    path,
			Keyword: "type",
			Message: `should be type "array"`,
		})

		return errors
	}

	if self.Items != nil {
		for i := 0; i < v.Len(); i++ {
			schema := self.Items.One

			if schema == nil && self.AdditionalItems != nil {
				schema = self.AdditionalItems.Schema
			}

			if self.Items.Many != nil && i < len(self.Items.Many) {
				schema = self.Items.Many[i]
			}

			if schema == nil {
				if self.AdditionalItems == nil || (self.AdditionalItems.Bool != nil && !*self.AdditionalItems.Bool) {
					errors = append(errors, SchemaError{
						Path:    fmt.Sprintf("%s/%s", path, strconv.Itoa(i)),
						Keyword: "additionalItems",
						Message: "undefined array index",
					})
				}
			}

			if schema != nil {
				errs := schema.validate(
					ns,
					path,
					strconv.Itoa(i),
					v.Index(i).Interface(),
				)

				if len(errs) > 0 {
					errors = append(errors, errs...)
				}
			}
		}
	}

	if self.MinItems != nil {
		if *self.MinItems > v.Len() {
			errors = append(errors, SchemaError{
				Path:    path,
				Keyword: "minItems",
				Message: fmt.Sprintf(
					`"%d" is not greater than or equal to "%d"`,
					v.Len(),
					*self.MinItems,
				),
			})
		}
	}

	if self.MaxItems != nil {
		if *self.MaxItems < v.Len() {
			errors = append(errors, SchemaError{
				Path:    path,
				Keyword: "maxItems",
				Message: fmt.Sprintf(
					`"%d" is not less than or equal to "%d"`,
					v.Len(),
					*self.MaxItems,
				),
			})
		}
	}

	if self.UniqueItems != nil && *self.UniqueItems {
		items := map[string]bool{}

		for i := 0; i < v.Len(); i++ {
			item, _ := json.Marshal(v.Index(i).Interface())
			_, ok := items[string(item)]

			if ok {
				errors = append(errors, SchemaError{
					Path:    fmt.Sprintf("%s/%d", path, i),
					Keyword: "uniqueItems",
					Message: "duplicate item",
				})
			}

			items[string(item)] = true
		}
	}

	return errors
}

func parseArray(data map[string]any) (ArraySchema, error) {
	self := ArraySchema{Type: SCHEMA_TYPE_ARRAY}

	if data == nil {
		return self, errors.New(`cannot parse "null" to "ArraySchema"`)
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

	if items, ok := data["items"]; ok {
		parsed, err := parseArrayItems(items)

		if err != nil {
			return self, err
		}

		self.Items = parsed
	}

	if additionalItems, ok := data["additionalItems"]; ok {
		parsed, err := parseArrayAdditionalItems(additionalItems)

		if err != nil {
			return self, nil
		}

		self.AdditionalItems = parsed
	}

	if minItems, ok := data["minItems"].(float64); ok {
		v := int(minItems)
		self.MinItems = &v
	}

	if maxItems, ok := data["maxItems"].(float64); ok {
		v := int(maxItems)
		self.MinItems = &v
	}

	if uniqueItems, ok := data["uniqueItems"].(bool); ok {
		self.UniqueItems = &uniqueItems
	}

	return self, nil
}
