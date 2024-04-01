package jsonschema_test

import (
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestComment(t *testing.T) {
	RunAll("./testcases/comment", jsonschema.New(), t)
}

func BenchmarkComment(b *testing.B) {
	RunAllBench("./testcases/comment", jsonschema.New(), b)
}
