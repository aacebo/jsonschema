package jsonschema

import (
	"fmt"
	"jsonschema/core"
)

type ArrayBuilder struct {
	schema *ArraySchema
}

func Array() *ArrayBuilder {
	return &ArrayBuilder{&ArraySchema{
		Type: core.SCHEMA_TYPE_ARRAY,
	}}
}

func (self *ArrayBuilder) ID(id string) *ArrayBuilder {
	self.schema.ID = &id
	return self
}

func (self *ArrayBuilder) Title(title string) *ArrayBuilder {
	self.schema.Title = &title
	return self
}

func (self *ArrayBuilder) Description(description string) *ArrayBuilder {
	self.schema.Description = &description
	return self
}

func (self *ArrayBuilder) Item(schema Schema) *ArrayBuilder {
	self.schema.Items = &ArrayItems{One: schema}
	return self
}

func (self *ArrayBuilder) Items(schemas ...Schema) *ArrayBuilder {
	self.schema.Items = &ArrayItems{Many: schemas}
	return self
}

func (self *ArrayBuilder) AdditionalItems(schema Schema) *ArrayBuilder {
	self.schema.AdditionalItems = &ArrayAdditionalItems{Schema: schema}
	return self
}

func (self *ArrayBuilder) AdditionalItemsAllowed(allow bool) *ArrayBuilder {
	self.schema.AdditionalItems = &ArrayAdditionalItems{Bool: &allow}
	return self
}

func (self *ArrayBuilder) MinItems(minItems int) *ArrayBuilder {
	self.schema.MinItems = &minItems
	return self
}

func (self *ArrayBuilder) MaxItems(maxItems int) *ArrayBuilder {
	self.schema.MaxItems = &maxItems
	return self
}

func (self *ArrayBuilder) UniqueItems(uniqueItems bool) *ArrayBuilder {
	self.schema.UniqueItems = &uniqueItems
	return self
}

func (self *ArrayBuilder) Build() *ArraySchema {
	errs := New().AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}

func (self *ArrayBuilder) BuildIn(ns core.Namespace[Schema]) *ArraySchema {
	errs := ns.AddSchema(self.schema).Compile(self.schema.GetID())

	if len(errs) > 0 {
		panic(fmt.Sprintf("failed to build with errors: %v", errs))
	}

	return self.schema
}
