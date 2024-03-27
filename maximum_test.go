package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMaximum(t *testing.T) {
	RunAll("./testcases/number/maximum", jsonschema.New(), t)
}
