package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestEnum(t *testing.T) {
	RunAll("./testcases/enum", jsonschema.New(), t)
}
