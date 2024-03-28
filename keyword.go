package jsonschema

type Context struct {
	Path   string
	Schema map[string]any
	Value  any
}

type Keyword struct {
	Default  any                                                       // the default configuration of the keyword
	Compile  func(ns *Namespace, ctx Context) []SchemaError            // used to validate the keywords configuration
	Validate func(ns *Namespace, ctx Context, input any) []SchemaError // used to validate a value using the keyword
}
