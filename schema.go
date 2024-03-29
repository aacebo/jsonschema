package jsonschema

import (
	"encoding/json"
)

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

func (self Schema) Spec() string {
	if v, ok := self["$schema"]; ok {
		return v.(string)
	}

	return "http://json-schema.org/draft-04/schema#"
}

func (self Schema) String() string {
	b, err := json.Marshal(self)

	if err != nil {
		return err.Error()
	}

	return string(b)
}
