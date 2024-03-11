package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/core"
	"math"
	"strconv"
)

// https://json-schema.org/understanding-json-schema/reference/numeric
type NumberSchema struct {
	ID               *string         `json:"$id,omitempty"`              // https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
	Type             core.SchemaType `json:"type"`                       // https://json-schema.org/understanding-json-schema/reference/type
	Title            *string         `json:"title,omitempty"`            // https://json-schema.org/understanding-json-schema/reference/annotations
	Description      *string         `json:"description,omitempty"`      // https://json-schema.org/understanding-json-schema/reference/annotations
	MultipleOf       *float64        `json:"multipleOf,omitempty"`       // https://json-schema.org/understanding-json-schema/reference/numeric#multiples
	Minimum          *float64        `json:"minimum,omitempty"`          // https://json-schema.org/understanding-json-schema/reference/numeric#range
	Maximum          *float64        `json:"maximum,omitempty"`          // https://json-schema.org/understanding-json-schema/reference/numeric#range
	ExclusiveMinimum *bool           `json:"exclusiveMinimum,omitempty"` // https://json-schema.org/understanding-json-schema/reference/numeric#range
	ExclusiveMaximum *bool           `json:"exclusiveMaximum,omitempty"` // https://json-schema.org/understanding-json-schema/reference/numeric#range
}

func (self NumberSchema) GetID() string {
	if self.ID != nil {
		return *self.ID
	}

	return ""
}

func (self NumberSchema) GetType() core.SchemaType {
	return self.Type
}

func (self NumberSchema) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self NumberSchema) compile(ns core.Namespace[Schema], id string, path string) []core.SchemaError {
	errors := []core.SchemaError{}

	if self.Type != core.SCHEMA_TYPE_NUMBER {
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "type",
			Message: fmt.Sprintf(`"type" must be "%s"`, core.SCHEMA_TYPE_NUMBER),
		})
	}

	if self.MultipleOf != nil {
		if *self.MultipleOf < 0 {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "multipleOf",
				Message: fmt.Sprintf("must be non-negative"),
			})
		}
	}

	if self.Minimum != nil {
		if *self.Minimum < 0 {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "minimum",
				Message: fmt.Sprintf("must be non-negative"),
			})
		}
	}

	if self.Maximum != nil {
		if *self.Maximum < 0 {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "maximum",
				Message: fmt.Sprintf("must be non-negative"),
			})
		}
	}

	if self.Minimum != nil && self.Maximum != nil && *self.Minimum > *self.Maximum {
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "maximum",
			Message: fmt.Sprintf(`must be greater than or equal to "minimum"`),
		})
	}

	return errors
}

func (self NumberSchema) validate(ns core.Namespace[Schema], id string, path string, value any) []core.SchemaError {
	var v float64
	errors := []core.SchemaError{}

	switch t := value.(type) {
	case int:
		v = float64(t)
		break
	case float32:
		v = float64(t)
		break
	case float64:
		v = t
		break
	default:
		errors = append(errors, core.SchemaError{
			Path:    path,
			Keyword: "type",
			Message: `should be type "number"`,
		})

		return errors
	}

	if self.MultipleOf != nil {
		if math.Mod(v, *self.MultipleOf) != 0 {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "multipleOf",
				Message: fmt.Sprintf(
					`"%s" is not a multiple of "%s"`,
					strconv.FormatFloat(v, 'f', -1, 64),
					strconv.FormatFloat(*self.MultipleOf, 'f', -1, 64),
				),
			})
		}
	}

	if self.Minimum != nil {
		min := *self.Minimum

		if self.ExclusiveMinimum != nil && *self.ExclusiveMinimum {
			min += 1
		}

		if v < min {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "minimum",
				Message: fmt.Sprintf(
					`"%s" is not greater than or equal to "%s"`,
					strconv.FormatFloat(v, 'f', -1, 64),
					strconv.FormatFloat(min, 'f', -1, 64),
				),
			})
		}
	}

	if self.Maximum != nil {
		max := *self.Maximum

		if self.ExclusiveMaximum != nil && *self.ExclusiveMaximum {
			max -= 1
		}

		if v > max {
			errors = append(errors, core.SchemaError{
				Path:    path,
				Keyword: "maximum",
				Message: fmt.Sprintf(
					`"%s" is not less than or equal to "%s"`,
					strconv.FormatFloat(v, 'f', -1, 64),
					strconv.FormatFloat(max, 'f', -1, 64),
				),
			})
		}
	}

	return errors
}

func parseNumber(data map[string]any) (NumberSchema, error) {
	self := NumberSchema{Type: core.SCHEMA_TYPE_NUMBER}

	if data == nil {
		return self, errors.New(`cannot parse "null" to "NumberSchema"`)
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

	if multipleOf, ok := data["multipleOf"].(float64); ok {
		self.MultipleOf = &multipleOf
	}

	if minimum, ok := data["minimum"].(float64); ok {
		self.Minimum = &minimum
	}

	if maximum, ok := data["maximum"].(float64); ok {
		self.Maximum = &maximum
	}

	if exclusiveMinimum, ok := data["exclusiveMinimum"].(bool); ok {
		self.ExclusiveMinimum = &exclusiveMinimum
	}

	if exclusiveMaximum, ok := data["exclusiveMaximum"].(bool); ok {
		self.ExclusiveMaximum = &exclusiveMaximum
	}

	return self, nil
}
