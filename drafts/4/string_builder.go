package jsonschema

import (
	"fmt"
	"jsonschema/core"
)

type StringBuilder struct {
	schema *StringSchema
}

func String() *StringBuilder {
	return &StringBuilder{&StringSchema{
		Type: core.SCHEMA_TYPE_STRING,
	}}
}

func (self *StringBuilder) ID(id string) *StringBuilder {
	self.schema.ID = &id
	return self
}

func (self *StringBuilder) Title(title string) *StringBuilder {
	self.schema.Title = &title
	return self
}

func (self *StringBuilder) Description(description string) *StringBuilder {
	self.schema.Description = &description
	return self
}

func (self *StringBuilder) Pattern(pattern string) *StringBuilder {
	self.schema.Pattern = &pattern
	return self
}

func (self *StringBuilder) Format(format string) *StringBuilder {
	self.schema.Format = &format
	return self
}

func (self *StringBuilder) MinLength(minLength int) *StringBuilder {
	self.schema.MinLength = &minLength
	return self
}

func (self *StringBuilder) MaxLength(maxLength int) *StringBuilder {
	self.schema.MaxLength = &maxLength
	return self
}

func (self *StringBuilder) Build() *StringSchema {
	errs := New().AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}

func (self *StringBuilder) BuildIn(ns core.Namespace[Schema]) *StringSchema {
	errs := ns.AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}
