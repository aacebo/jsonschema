package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestComment(t *testing.T) {
	RunAll("./testcases/comment", jsonschema.New(), t)
}
