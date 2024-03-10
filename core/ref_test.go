package core_test

import (
	"jsonschema/core"
	"testing"
)

func TestRef(t *testing.T) {
	// https://json-schema.org/understanding-json-schema/structuring#defs
	t.Run("should resolve", func(t *testing.T) {
		url, err := core.ResolveRef(
			"https://example.com/schemas/customer",
			"/schemas/address",
		)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if url != "https://example.com/schemas/address" {
			t.Log(url)
			t.FailNow()
		}
	})

	// https://json-schema.org/understanding-json-schema/structuring#recursion
	t.Run("should resolve recursive", func(t *testing.T) {
		url, err := core.ResolveRef(
			"https://example.com/schemas/customer",
			"#",
		)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if url != "https://example.com/schemas/customer" {
			t.Log(url)
			t.FailNow()
		}
	})

	// https://json-schema.org/understanding-json-schema/structuring#defs
	t.Run("should resolve defs", func(t *testing.T) {
		url, err := core.ResolveRef(
			"",
			"#/$defs/test",
		)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if url != "#/$defs/test" {
			t.Log(url)
			t.FailNow()
		}
	})

	// https://json-schema.org/understanding-json-schema/structuring#json-pointer
	t.Run("should resolve json pointer", func(t *testing.T) {
		url, err := core.ResolveRef(
			"https://example.com/schemas/customer",
			"#/properties/test",
		)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if url != "https://example.com/schemas/customer#/properties/test" {
			t.Log(url)
			t.FailNow()
		}
	})
}
