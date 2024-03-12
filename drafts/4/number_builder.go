package jsonschema

import (
	"fmt"
	"jsonschema/core"
)

type NumberBuilder struct {
	schema *NumberSchema
}

func Number() *NumberBuilder {
	return &NumberBuilder{&NumberSchema{
		Type: core.SCHEMA_TYPE_NUMBER,
	}}
}

func (self *NumberBuilder) ID(id string) *NumberBuilder {
	self.schema.ID = &id
	return self
}

func (self *NumberBuilder) Title(title string) *NumberBuilder {
	self.schema.Title = &title
	return self
}

func (self *NumberBuilder) Description(description string) *NumberBuilder {
	self.schema.Description = &description
	return self
}

func (self *NumberBuilder) MultipleOf(multipleOf float64) *NumberBuilder {
	self.schema.MultipleOf = &multipleOf
	return self
}

func (self *NumberBuilder) Minimum(minimum float64) *NumberBuilder {
	self.schema.Minimum = &minimum
	return self
}

func (self *NumberBuilder) Maximum(maximum float64) *NumberBuilder {
	self.schema.Maximum = &maximum
	return self
}

func (self *NumberBuilder) Build() *NumberSchema {
	errs := New().AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}

func (self *NumberBuilder) BuildIn(ns core.Namespace[Schema]) *NumberSchema {
	errs := ns.AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}
