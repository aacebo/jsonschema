package jsonschema

import "reflect"

type Context struct {
	Path   string
	Schema Schema
}

type Keyword struct {
	Default  any                                                                                       // the default configuration of the keyword
	Compile  func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError                      // used to validate the keywords configuration
	Validate func(ns *Namespace, ctx Context, config reflect.Value, value reflect.Value) []SchemaError // used to validate a value using the keyword
}
