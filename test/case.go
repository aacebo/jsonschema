package test

import (
	"jsonschema/core"
)

type Case struct {
	Schema core.Schema
	Input  any
	Errors []string
}
