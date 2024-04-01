package jsonschema_test

import (
	"jsonschema"
	"testing"
)

func TestSchemaBuilder(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		ns := jsonschema.New()

		t.Run("should succeed", func(t *testing.T) {
			schema := jsonschema.Builder().
				String().
				Pattern("^[0-9]*$").
				Build()

			errs := ns.Compile(schema)

			if len(errs) > 0 {
				t.Error(errs)
			}

			errs = ns.Validate(schema, "123")

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})

	t.Run("integer", func(t *testing.T) {
		ns := jsonschema.New()

		t.Run("should succeed", func(t *testing.T) {
			schema := jsonschema.Builder().
				Integer().
				Minimum(20).
				Build()

			errs := ns.Compile(schema)

			if len(errs) > 0 {
				t.Error(errs)
			}

			errs = ns.Validate(schema, 21)

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})
}
