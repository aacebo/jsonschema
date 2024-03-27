package jsonschema

type Context struct {
	Path   string
	Schema map[string]any
	Value  any
}

type Keyword struct {
	Compile  func(ns *Namespace, ctx Context) []SchemaError
	Validate func(ns *Namespace, ctx Context, input any) []SchemaError
}
