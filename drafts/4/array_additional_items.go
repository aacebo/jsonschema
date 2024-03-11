package jsonschema

import (
	"encoding/json"
	"errors"
)

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
type ArrayAdditionalItems struct {
	Bool   *bool
	Schema Schema
}

func (self ArrayAdditionalItems) Value() any {
	if self.Bool != nil {
		return *self.Bool
	}

	return self.Schema.Value()
}

func (self ArrayAdditionalItems) MarshalJSON() ([]byte, error) {
	var data []byte

	if self.Bool != nil {
		b, err := json.Marshal(self.Bool)

		if err != nil {
			return nil, err
		}

		data = b
	}

	if self.Schema != nil {
		b, err := json.Marshal(self.Schema)

		if err != nil {
			return nil, err
		}

		data = b
	}

	return data, nil
}

func (self *ArrayAdditionalItems) UnmarshalJSON(data []byte) error {
	var value any
	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	items, err := parseArrayAdditionalItems(value)

	if err != nil {
		return err
	}

	*self = *items
	return nil
}

func parseArrayAdditionalItems(data any) (*ArrayAdditionalItems, error) {
	self := ArrayAdditionalItems{}

	switch v := data.(type) {
	case bool:
		self.Bool = &v
		break
	case map[string]any:
		schema, err := parse(v)

		if err != nil {
			return nil, err
		}

		self.Schema = schema
		break
	default:
		return nil, errors.New(`"array.additionalItems" must be either a "bool" or "Schema"`)
	}

	return &self, nil
}
