package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestRef(t *testing.T) {
	t.Run("result", func(t *testing.T) {
		// https://json-schema.org/understanding-json-schema/structuring#defs
		t.Run("should resolve", func(t *testing.T) {
			ref, err := jsonschema.RefSchema{"/schemas/address"}.Parse(
				"https://example.com/schemas/customer",
			)

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if ref.String() != "https://example.com/schemas/address" {
				t.Log(ref)
				t.FailNow()
			}
		})

		// https://json-schema.org/understanding-json-schema/structuring#recursion
		t.Run("should resolve recursive", func(t *testing.T) {
			ref, err := jsonschema.RefSchema{"#"}.Parse(
				"https://example.com/schemas/customer",
			)

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if ref.String() != "https://example.com/schemas/customer" {
				t.Log(ref)
				t.FailNow()
			}
		})

		// https://json-schema.org/understanding-json-schema/structuring#defs
		t.Run("should resolve defs", func(t *testing.T) {
			ref, err := jsonschema.RefSchema{"#/$defs/test"}.Parse(
				"https://example.com/schemas/customer",
			)

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if ref.String() != "https://example.com/schemas/customer#/$defs/test" {
				t.Log(ref)
				t.FailNow()
			}
		})

		// https://json-schema.org/understanding-json-schema/structuring#json-pointer
		t.Run("should resolve json pointer", func(t *testing.T) {
			ref, err := jsonschema.RefSchema{"#/properties/test"}.Parse(
				"https://example.com/schemas/customer",
			)

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if ref.String() != "https://example.com/schemas/customer#/properties/test" {
				t.Log(ref)
				t.FailNow()
			}
		})

		t.Run("should resolve non url", func(t *testing.T) {
			ref, err := jsonschema.RefSchema{"#/properties/test"}.Parse(
				"root",
			)

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if ref.String() != "root#/properties/test" {
				t.Log(ref)
				t.FailNow()
			}

			if ref.URL != "root" {
				t.Log(ref.URL)
				t.FailNow()
			}

			if ref.Path != "#/properties/test" {
				t.Log(ref.Path)
				t.FailNow()
			}
		})
	})

	t.Run("schema", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			ns := jsonschema.New()
			_, err := ns.Read("./schemas/array/ref.json")

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			errs := ns.Compile("root")

			if len(errs) > 0 {
				t.Log(errs)
				t.FailNow()
			}

			errs = ns.Validate("root", []any{"test"})

			if len(errs) > 0 {
				t.Log(errs)
				t.FailNow()
			}
		})

		t.Run("should error on type", func(t *testing.T) {
			ns := jsonschema.New()
			_, err := ns.Read("./schemas/array/ref.json")

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			errs := ns.Compile("root")

			if len(errs) > 0 {
				t.Log(errs)
				t.FailNow()
			}

			errs = ns.Validate("root", []any{"test", 1})

			if len(errs) == 0 {
				t.Log("should have errors")
				t.FailNow()
			}
		})
	})
}
