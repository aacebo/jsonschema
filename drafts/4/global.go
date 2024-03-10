package jsonschema

var global = New()

func Read(path string) (Schema, error) {
	return global.Read(path)
}

func HasSchema(id string) bool {
	return global.HasSchema(id)
}

func GetSchema(id string) Schema {
	return global.GetSchema(id)
}

func AddSchema(schema Schema) {
	global.AddSchema(schema)
}

func Compile(id string) []SchemaCompileError {
	return global.Compile(id)
}

func Validate(id string, value any) []SchemaError {
	return global.Validate(id, value)
}
