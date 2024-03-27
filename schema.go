package jsonschema

import (
	"encoding/json"
)

// type Schema struct {
// 	ID           *ID          `json:"$id,omitempty"`          // https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
// 	Schema       *SchemaSpec  `json:"$schema,omitempty"`      // https://json-schema.org/understanding-json-schema/reference/schema#declaring-a-dialect
// 	Type         []SchemaType `json:"type,omitempty"`         // https://json-schema.org/understanding-json-schema/reference/type
// 	Title        *Title       `json:"title,omitempty"`        // https://json-schema.org/understanding-json-schema/reference/annotations
// 	Description  *Description `json:"description,omitempty"`  // https://json-schema.org/understanding-json-schema/reference/annotations
// 	Dependencies Dependencies `json:"dependencies,omitempty"` // https://json-schema.org/understanding-json-schema/reference/conditionals#dependentRequired

// 	// string
// 	Pattern   *Pattern   `json:"pattern,omitempty"`   // https://json-schema.org/understanding-json-schema/reference/string#regexp
// 	Format    *Format    `json:"format,omitempty"`    // https://json-schema.org/understanding-json-schema/reference/string#format
// 	MinLength *MinLength `json:"minLength,omitempty"` // https://json-schema.org/understanding-json-schema/reference/string#length
// 	MaxLength *MaxLength `json:"maxLength,omitempty"` // https://json-schema.org/understanding-json-schema/reference/string#length

// 	// number
// 	MultipleOf *MultipleOf `json:"multipleOf,omitempty"` // https://json-schema.org/understanding-json-schema/reference/numeric#multiples
// 	Minimum    *Minimum    `json:"minimum,omitempty"`    // https://json-schema.org/understanding-json-schema/reference/numeric#range
// 	Maximum    *Maximum    `json:"maximum,omitempty"`    // https://json-schema.org/understanding-json-schema/reference/numeric#range
// 	// ExclusiveMinimum *float64    `json:"exclusiveMinimum,omitempty"` // https://json-schema.org/understanding-json-schema/reference/numeric#range
// 	// ExclusiveMaximum *float64    `json:"exclusiveMaximum,omitempty"` // https://json-schema.org/understanding-json-schema/reference/numeric#range
// }

type Schema map[string]any

func (self Schema) ID() string {
	if v, ok := self["$id"]; ok {
		return v.(string)
	}

	if v, ok := self["id"]; ok {
		return v.(string)
	}

	return ""
}

func (self Schema) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
