package test

import (
	"jsonschema/core"
)

type Case struct {
	Schema core.Schema
	Input  string
	Errors []string
}
