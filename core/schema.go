package core

type SchemaType string

const (
	SCHEMA_TYPE_STRING  SchemaType = "string"
	SCHEMA_TYPE_NUMBER  SchemaType = "number"
	SCHEMA_TYPE_NULL    SchemaType = "null"
	SCHEMA_TYPE_INTEGER SchemaType = "integer"
	SCHEMA_TYPE_BOOLEAN SchemaType = "boolean"
	SCHEMA_TYPE_ARRAY   SchemaType = "array"
	SCHEMA_TYPE_OBJECT  SchemaType = "object"
	SCHEMA_TYPE_REF     SchemaType = "reference"
)

type Schema interface {
	GetID() string
	GetType() SchemaType
}
