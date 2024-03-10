package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
)

// https://json-schema.org/understanding-json-schema/reference/array#items
type ArrayItems struct {
	One  Schema
	Many []Schema
}

func (self ArrayItems) MarshalJSON() ([]byte, error) {
	var data []byte

	if self.One != nil {
		b, err := json.Marshal(self.One)

		if err != nil {
			return nil, err
		}

		data = b
	}

	if self.Many != nil {
		b, err := json.Marshal(self.Many)

		if err != nil {
			return nil, err
		}

		data = b
	}

	return data, nil
}

func (self *ArrayItems) UnmarshalJSON(data []byte) error {
	var value any
	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	items, err := parseArrayItems(value)

	if err != nil {
		return err
	}

	*self = *items
	return nil
}

func parseArrayItems(data any) (*ArrayItems, error) {
	self := ArrayItems{}

	switch v := data.(type) {
	case map[string]any:
		one, err := parse(v)

		if err != nil {
			return nil, err
		}

		self.One = one
		break
	case []any:
		many := []Schema{}

		for i, item := range v {
			m, ok := item.(map[string]any)

			if !ok {
				return nil, SchemaCompileError{
					Path:    fmt.Sprintf("items/%d", i),
					Message: `must be a "Schema"`,
				}
			}

			schema, err := parse(m)

			if err != nil {
				return nil, err
			}

			many = append(many, schema)
		}

		self.Many = many
		break
	default:
		return nil, errors.New(`"array.items" must be either a "Schema" or "[]Schema"`)
	}

	return &self, nil
}
