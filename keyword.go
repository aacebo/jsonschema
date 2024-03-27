package jsonschema

type Context struct {
	Path   string
	Schema map[string]any
	Value  any
}

type Keyword struct {
	compile  func(ns *Namespace, ctx Context) []SchemaError
	validate func(ns *Namespace, ctx Context, input any) []SchemaError
}
