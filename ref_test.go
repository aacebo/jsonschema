package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestRef(t *testing.T) {
	RunAll("./testcases/ref", jsonschema.New(), t)
}
