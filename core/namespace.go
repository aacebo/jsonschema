package core

type Namespace interface {
	HasFormat(name string) bool
	AddFormat(name string, format Formatter) Namespace
	Format(name string, input string) error

	HasSchema(id string) bool
	GetSchema(id string) Schema
	AddSchema(schema Schema) Namespace

	Read(path string) (Schema, error)
	Compile(id string) []SchemaError
	Validate(id string, value any) []SchemaError
}
