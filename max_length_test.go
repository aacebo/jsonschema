package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaxLength(t *testing.T) {
	RunAll("./testcases/string/max_length", jsonschema.New(), t)
}