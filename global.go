package jsonschema

var namespace = New()

func HasFormat(name string) bool {
	return namespace.HasFormat(name)
}

func AddFormat(name string, format Formatter) {
	namespace.AddFormat(name, format)
}

func AddKeyword(name string, keyword Keyword) {
	namespace.AddKeyword(name, keyword)
}

func Compile(schema Schema) []SchemaError {
	return namespace.Compile(schema)
}

func Validate(schema Schema, value any) []SchemaError {
	return namespace.Validate(schema, value)
}

func Read(path string) (Schema, error) {
	return namespace.Read(path)
}
