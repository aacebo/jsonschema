package jsonschema

type SchemaBuilder struct {
	schema Schema
}

func NewSchema() *SchemaBuilder {
	return &SchemaBuilder{Schema{}}
}

func (self *SchemaBuilder) ID(id string) *SchemaBuilder {
	self.schema["$id"] = id
	return self
}

func (self *SchemaBuilder) Schema(schema string) *SchemaBuilder {
	self.schema["$schema"] = schema
	return self
}

func (self *SchemaBuilder) Title(title string) *SchemaBuilder {
	self.schema["title"] = title
	return self
}

func (self *SchemaBuilder) Description(description string) *SchemaBuilder {
	self.schema["description"] = description
	return self
}

func (self *SchemaBuilder) Type(types ...SchemaType) *SchemaBuilder {
	if len(types) == 0 {
		return self
	}

	if len(types) == 1 {
		self.schema["type"] = types[0]
	} else {
		self.schema["type"] = types
	}

	return self
}

func (self *SchemaBuilder) Build() Schema {
	return self.schema
}
