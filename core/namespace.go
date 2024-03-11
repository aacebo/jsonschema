package core

type Namespace[T Schema] interface {
	HasFormat(name string) bool
	AddFormat(name string, format Formatter) Namespace[T]
	Format(name string, input string) error

	HasSchema(id string) bool
	GetSchema(id string) T
	AddSchema(schema T) Namespace[T]

	Read(path string) (T, error)
	Resolve(url string, path string) (T, error)
	Compile(id string) []SchemaError
	Validate(id string, value any) []SchemaError
}
