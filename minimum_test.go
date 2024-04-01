package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestMinimum(t *testing.T) {
	RunAll("./testcases/number/minimum", jsonschema.New(), t)
	RunAll("./testcases/integer/minimum", jsonschema.New(), t)
}
