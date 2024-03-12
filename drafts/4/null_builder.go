package jsonschema

import (
	"fmt"
	"jsonschema/core"
)

type NullBuilder struct {
	schema *NullSchema
}

func Null() *NullBuilder {
	return &NullBuilder{&NullSchema{
		Type: core.SCHEMA_TYPE_NULL,
	}}
}

func (self *NullBuilder) ID(id string) *NullBuilder {
	self.schema.ID = &id
	return self
}

func (self *NullBuilder) Title(title string) *NullBuilder {
	self.schema.Title = &title
	return self
}

func (self *NullBuilder) Description(description string) *NullBuilder {
	self.schema.Description = &description
	return self
}

func (self *NullBuilder) Build() *NullSchema {
	errs := New().AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}

func (self *NullBuilder) BuildIn(ns core.Namespace[Schema]) *NullSchema {
	errs := ns.AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}
