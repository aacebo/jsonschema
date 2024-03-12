package jsonschema

import (
	"jsonschema/test"
	"testing"
)

func TestGlobal(t *testing.T) {
	test.RunAll("./testcases", global, t)
}
