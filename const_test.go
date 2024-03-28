package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestConst(t *testing.T) {
	RunAll("./testcases/const", jsonschema.New(), t)
}
