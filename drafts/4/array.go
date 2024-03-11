package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/core"
	"reflect"
)

// https://json-schema.org/understanding-json-schema/reference/array
type ArraySchema struct {
	ID              *string               `json:"$id,omitempty"`             // https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
	Type            core.SchemaType       `json:"type"`                      // https://json-schema.org/understanding-json-schema/reference/type
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

func (self ArraySchema) GetType() core.SchemaType {
	return self.Type
}

func (self ArraySchema) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self ArraySchema) compile(ns core.Namespace[Schema], id string, path string) []core.SchemaError {
	errors := []core.SchemaError{}

	if self.ID != nil && *self.ID != "" {
		id = *self.ID
	}

	if self.Type != core.SCHEMA_TYPE_ARRAY {
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "type",
			Message: fmt.Sprintf(`"type" must be "%s"`, core.SCHEMA_TYPE_ARRAY),
		})
	}

	if self.Items != nil {
		if self.Items.One != nil {
			if errs := self.Items.One.compile(ns, id, path+"/items"); len(errs) > 0 {
				errors = append(errors, errs...)
			}
		}

		if self.Items.Many != nil {
			for i, schema := range self.Items.Many {
				if errs := schema.compile(ns, id, fmt.Sprintf("%s/items/%d", path, i)); len(errs) > 0 {
					errors = append(errors, errs...)
				}
			}
		}
	}

	if self.AdditionalItems != nil {
		if self.AdditionalItems.Schema != nil {
			errs := self.AdditionalItems.Schema.compile(ns, id, path+"/additionalItems")

			if len(errs) > 0 {
				errors = append(errors, errs...)
			}
		}
	}

	if self.MinItems != nil {
		if *self.MinItems < 0 {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "minItems",
				Message: "must be non-negative",
			})
		}
	}

	if self.MaxItems != nil {
		if *self.MaxItems < 0 {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "maxItems",
				Message: "must be non-negative",
			})
		}
	}

	if self.MinItems != nil && self.MaxItems != nil && *self.MaxItems < *self.MinItems {
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "maxItems",
			Message: `must be greater than or equal to "minItems"`,
		})
	}

	return errors
}

func (self ArraySchema) validate(ns core.Namespace[Schema], id string, path string, value any) []core.SchemaError {
	errors := []core.SchemaError{}
	v := reflect.ValueOf(value)

	if self.ID != nil && *self.ID != "" {
		id = *self.ID
	}

	if v.Kind() != reflect.Slice {
		errors = append(errors, core.SchemaError{
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
					errors = append(errors, core.SchemaError{
						Path:    fmt.Sprintf("%s/items/%d", path, i),
						Keyword: "additionalItems",
						Message: "undefined array index",
					})
				}
			}

			if schema != nil {
				errs := schema.validate(
					ns,
					id,
					fmt.Sprintf("%s/items/%d", path, i),
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
			errors = append(errors, core.SchemaError{
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
			errors = append(errors, core.SchemaError{
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
				errors = append(errors, core.SchemaError{
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
	self := ArraySchema{Type: core.SCHEMA_TYPE_ARRAY}

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
