package jsonschema

type SchemaBuilder struct {
	schema Schema
}

func Builder() *SchemaBuilder {
	return &SchemaBuilder{Schema{}}
}

// https://json-schema.org/understanding-json-schema/basics#declaring-a-unique-identifier
func (self *SchemaBuilder) ID(id string) *SchemaBuilder {
	self.schema["$id"] = id
	return self
}

// https://json-schema.org/understanding-json-schema/reference/schema#declaring-a-dialect
func (self *SchemaBuilder) Schema(schema string) *SchemaBuilder {
	self.schema["$schema"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/comments#comments
func (self *SchemaBuilder) Comment(comment string) *SchemaBuilder {
	self.schema["$comment"] = comment
	return self
}

// https://json-schema.org/understanding-json-schema/structuring#dollarref
func (self *SchemaBuilder) Ref(ref string) *SchemaBuilder {
	self.schema["$ref"] = ref
	return self
}

// https://json-schema.org/understanding-json-schema/reference/annotations
func (self *SchemaBuilder) Title(title string) *SchemaBuilder {
	self.schema["title"] = title
	return self
}

// https://json-schema.org/understanding-json-schema/reference/annotations
func (self *SchemaBuilder) Description(description string) *SchemaBuilder {
	self.schema["description"] = description
	return self
}

// https://json-schema.org/understanding-json-schema/reference/annotations#annotations
func (self *SchemaBuilder) Default(defaultValue any) *SchemaBuilder {
	self.schema["default"] = defaultValue
	return self
}

// https://json-schema.org/understanding-json-schema/reference/const#constant-values
func (self *SchemaBuilder) Const(constValue any) *SchemaBuilder {
	self.schema["const"] = constValue
	return self
}

// https://json-schema.org/understanding-json-schema/reference/enum#enumerated-values
func (self *SchemaBuilder) Enum(options ...any) *SchemaBuilder {
	self.schema["enum"] = options
	return self
}

// https://json-schema.org/understanding-json-schema/structuring#defs
func (self *SchemaBuilder) Definitions(definitions map[string][]string) *SchemaBuilder {
	self.schema["definitions"] = definitions
	return self
}

// https://json-schema.org/understanding-json-schema/reference/combining#allOf
func (self *SchemaBuilder) AllOf(schemas ...Schema) *SchemaBuilder {
	self.schema["allOf"] = schemas
	return self
}

// https://json-schema.org/understanding-json-schema/reference/combining#anyOf
func (self *SchemaBuilder) AnyOf(schemas ...Schema) *SchemaBuilder {
	self.schema["anyOf"] = schemas
	return self
}

// https://json-schema.org/understanding-json-schema/reference/combining#oneOf
func (self *SchemaBuilder) OneOf(schemas ...Schema) *SchemaBuilder {
	self.schema["oneOf"] = schemas
	return self
}

// https://json-schema.org/understanding-json-schema/reference/combining#not
func (self *SchemaBuilder) Not(schema Schema) *SchemaBuilder {
	self.schema["not"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/type
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

// https://json-schema.org/understanding-json-schema/reference/string
func (self *SchemaBuilder) String() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_STRING
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#integer
func (self *SchemaBuilder) Integer() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_INTEGER
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#number
func (self *SchemaBuilder) Number() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_NUMBER
	return self
}

// https://json-schema.org/understanding-json-schema/reference/null
func (self *SchemaBuilder) Null() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_NULL
	return self
}

// https://json-schema.org/understanding-json-schema/reference/boolean
func (self *SchemaBuilder) Boolean() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_BOOLEAN
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array
func (self *SchemaBuilder) Array() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_ARRAY
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object
func (self *SchemaBuilder) Object() *SchemaBuilder {
	self.schema["type"] = SCHEMA_TYPE_OBJECT
	return self
}

// string

// https://json-schema.org/understanding-json-schema/reference/string#regexp
func (self *SchemaBuilder) Pattern(pattern string) *SchemaBuilder {
	self.schema["pattern"] = pattern
	return self
}

// https://json-schema.org/understanding-json-schema/reference/string#format
func (self *SchemaBuilder) Format(format string) *SchemaBuilder {
	self.schema["format"] = format
	return self
}

// https://json-schema.org/understanding-json-schema/reference/string#length
func (self *SchemaBuilder) MinLength(minLength int) *SchemaBuilder {
	self.schema["minLength"] = minLength
	return self
}

// https://json-schema.org/understanding-json-schema/reference/string#length
func (self *SchemaBuilder) MaxLength(maxLength int) *SchemaBuilder {
	self.schema["maxLength"] = maxLength
	return self
}

// number

// https://json-schema.org/understanding-json-schema/reference/numeric#multiples
func (self *SchemaBuilder) MultipleOf(multipleOf float64) *SchemaBuilder {
	self.schema["multipleOf"] = multipleOf
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) Minimum(minimum float64) *SchemaBuilder {
	self.schema["minimum"] = minimum
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) Maximum(maximum float64) *SchemaBuilder {
	self.schema["maximum"] = maximum
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) ExclusiveMinimum(exclusiveMinimum float64) *SchemaBuilder {
	self.schema["exclusiveMinimum"] = exclusiveMinimum
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) ExclusiveMinimumEnable() *SchemaBuilder {
	self.schema["exclusiveMinimum"] = true
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) ExclusiveMinimumDisable() *SchemaBuilder {
	self.schema["exclusiveMinimum"] = false
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) ExclusiveMaximum(exclusiveMaximum float64) *SchemaBuilder {
	self.schema["exclusiveMaximum"] = exclusiveMaximum
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) ExclusiveMaximumEnable() *SchemaBuilder {
	self.schema["exclusiveMaximum"] = true
	return self
}

// https://json-schema.org/understanding-json-schema/reference/numeric#range
func (self *SchemaBuilder) ExclusiveMaximumDisable() *SchemaBuilder {
	self.schema["exclusiveMaximum"] = false
	return self
}

// array

// https://json-schema.org/understanding-json-schema/reference/array#items
func (self *SchemaBuilder) Items(schema Schema) *SchemaBuilder {
	self.schema["items"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#items
func (self *SchemaBuilder) TupleItems(schemas ...Schema) *SchemaBuilder {
	self.schema["items"] = schemas
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
func (self *SchemaBuilder) AdditionalItems(schema Schema) *SchemaBuilder {
	self.schema["additionalItems"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
func (self *SchemaBuilder) AdditionalItemsEnable() *SchemaBuilder {
	self.schema["additionalItems"] = true
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#additionalitems
func (self *SchemaBuilder) AdditionalItemsDisable() *SchemaBuilder {
	self.schema["additionalItems"] = false
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#contains
func (self *SchemaBuilder) Contains(schema Schema) *SchemaBuilder {
	self.schema["contains"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#length
func (self *SchemaBuilder) MinItems(minItems int) *SchemaBuilder {
	self.schema["minItems"] = minItems
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#length
func (self *SchemaBuilder) MaxItems(maxItems int) *SchemaBuilder {
	self.schema["maxItems"] = maxItems
	return self
}

// https://json-schema.org/understanding-json-schema/reference/array#uniqueItems
func (self *SchemaBuilder) UniqueItems() *SchemaBuilder {
	self.schema["uniqueItems"] = true
	return self
}

// object

// https://json-schema.org/understanding-json-schema/reference/object#properties
func (self *SchemaBuilder) Properties(properties map[string]Schema) *SchemaBuilder {
	self.schema["properties"] = properties
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#patternProperties
func (self *SchemaBuilder) PatternProperties(patternProperties map[string]Schema) *SchemaBuilder {
	self.schema["patternProperties"] = patternProperties
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#additionalproperties
func (self *SchemaBuilder) AdditionalProperties(schema Schema) *SchemaBuilder {
	self.schema["additionalProperties"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#additionalproperties
func (self *SchemaBuilder) AdditionalPropertiesDisable() *SchemaBuilder {
	self.schema["additionalProperties"] = false
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#additionalproperties
func (self *SchemaBuilder) AdditionalPropertiesEnable() *SchemaBuilder {
	self.schema["additionalProperties"] = true
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#propertyNames
func (self *SchemaBuilder) PropertyNames(schema Schema) *SchemaBuilder {
	self.schema["propertyNames"] = schema
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#length
func (self *SchemaBuilder) MinProperties(minProperties int) *SchemaBuilder {
	self.schema["minProperties"] = minProperties
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#length
func (self *SchemaBuilder) MaxProperties(maxProperties int) *SchemaBuilder {
	self.schema["maxProperties"] = maxProperties
	return self
}

// https://json-schema.org/understanding-json-schema/reference/object#required
func (self *SchemaBuilder) Required(required ...string) *SchemaBuilder {
	self.schema["required"] = required
	return self
}

// other

func (self *SchemaBuilder) Build() Schema {
	return self.schema
}
