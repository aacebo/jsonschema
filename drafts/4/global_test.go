package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestGlobal(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should error on invalid type", func(t *testing.T) {
				schema, err := jsonschema.Read("./schemas/invalid/invalid_type.json")

				if err == nil || err.Error() != `invalid schema type "test"` {
					t.Log("should have error")
					t.FailNow()
				}

				if schema != nil {
					t.Logf(`"%s" should be nil`, schema)
					t.FailNow()
				}
			})

			t.Run("should error on missing type", func(t *testing.T) {
				schema, err := jsonschema.Read("./schemas/invalid/missing_type.json")

				if err == nil || err.Error() != "schema type is required and must be a string" {
					t.Log("should have error")
					t.FailNow()
				}

				if schema != nil {
					t.Logf(`"%s" should be nil`, schema)
					t.FailNow()
				}
			})
		})

		t.Run("string", func(t *testing.T) {
			t.Run("should parse", func(t *testing.T) {
				s, err := jsonschema.Read("./schemas/string/numeric.json")

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				schema, ok := s.(jsonschema.StringSchema)

				if !ok {
					t.Logf(`schema "%s" is not a "StringSchema"`, s)
					t.FailNow()
				}

				errs := jsonschema.Compile(schema.GetID())

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}

				if schema.ID == nil || *schema.ID != "root" {
					t.Log(`"$id" should be "root"`)
					t.FailNow()
				}

				if schema.Title == nil || *schema.Title != "only numbers" {
					t.Log(`"title" should be "only numbers"`)
					t.FailNow()
				}

				if schema.Description == nil || *schema.Description != "only numbers allowed" {
					t.Log(`"description" should be "only numbers allowed"`)
					t.FailNow()
				}

				if schema.Pattern == nil || *schema.Pattern != "^[0-9]*$" {
					t.Log(`"pattern" should be "^[0-9]*$"`)
					t.FailNow()
				}

				if schema.MinLength == nil || *schema.MinLength != 0 {
					t.Log(`"minLength" should be 0`)
					t.FailNow()
				}

				if schema.MaxLength == nil || *schema.MaxLength != 100 {
					t.Log(`"maxLength" should be 100`)
					t.FailNow()
				}
			})
		})

		t.Run("number", func(t *testing.T) {
			t.Run("should parse", func(t *testing.T) {
				s, err := jsonschema.Read("./schemas/number/range.json")

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				schema, ok := s.(jsonschema.NumberSchema)

				if !ok {
					t.Logf(`schema "%s" is not a "NumberSchema"`, s)
					t.FailNow()
				}

				errs := jsonschema.Compile(schema.GetID())

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}

				if schema.ID == nil || *schema.ID != "root" {
					t.Log(`"$id" should be "root"`)
					t.FailNow()
				}

				if schema.Title == nil || *schema.Title != "range" {
					t.Log(`"title" should be "range"`)
					t.FailNow()
				}

				if schema.Description == nil || *schema.Description != "bounded number range" {
					t.Log(`"description" should be "bounded number range"`)
					t.FailNow()
				}

				if schema.MultipleOf == nil || *schema.MultipleOf != 10 {
					t.Log(`"multipleOf" should be 10`)
					t.FailNow()
				}

				if schema.Minimum == nil || *schema.Minimum != 0 {
					t.Log(`"minimum" should be 0`)
					t.FailNow()
				}

				if schema.Maximum == nil || *schema.Maximum != 100 {
					t.Log(`"maximum" should be 100`)
					t.FailNow()
				}
			})
		})
	})
}
