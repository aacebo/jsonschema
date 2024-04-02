package jsonschema_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aacebo/jsonschema"
)

func TestFormat(t *testing.T) {
	RunAll("./testcases/string/format", jsonschema.New(), t)

	t.Run("should succeed", func(t *testing.T) {
		jsonschema.AddFormat("lowercase", func(input string) error {
			if strings.ToLower(input) != input {
				return errors.New("must be lowercase")
			}

			return nil
		})

		schema := jsonschema.Builder().
			String().
			Format("lowercase").
			Build()

		errs := jsonschema.Compile(schema)

		if len(errs) > 0 {
			t.Error(errs)
		}

		errs = jsonschema.Validate(schema, "test")

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func BenchmarkFormat(b *testing.B) {
	RunAllBench("./testcases/string/format", jsonschema.New(), b)
}
