package jsonschema_test

import (
	"encoding/json"
	"jsonschema/core"
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("$id", func(t *testing.T) {
		schema := jsonschema.ArraySchema{}
		err := json.Unmarshal([]byte(`{"$id": "123", "type": "array"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.GetID() != "123" {
			t.FailNow()
		}
	})

	t.Run("type", func(t *testing.T) {
		schema := jsonschema.ArraySchema{}
		err := json.Unmarshal([]byte(`{"type": "array"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.GetType() != core.SCHEMA_TYPE_ARRAY {
			t.Logf(`expected "number", received "%s"`, schema.GetType())
			t.FailNow()
		}

		if schema.GetID() != "" {
			t.Logf(`expected "", received "%s"`, schema.GetID())
			t.FailNow()
		}
	})

	t.Run("string", func(t *testing.T) {
		schema := jsonschema.ArraySchema{}
		err := json.Unmarshal([]byte(`{"type": "array", "title": "hello world"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.String() != `{"type":"array","title":"hello world"}` {
			t.Log(schema)
			t.FailNow()
		}
	})

	t.Run("compile", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				if schema.Title != nil {
					t.FailNow()
				}

				if schema.Description != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "title": "hello world"}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/type] => "type" must be "array"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})

		t.Run("should succeed with title", func(t *testing.T) {
			schema := jsonschema.ArraySchema{}
			err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "title": "hello world"}`), &schema)

			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if schema.Title == nil || *schema.Title != "hello world" {
				t.FailNow()
			}

			if schema.Description != nil {
				t.FailNow()
			}

			errs := jsonschema.New().AddSchema(schema).Compile("test")

			if len(errs) > 0 {
				t.Log(errs)
				t.FailNow()
			}
		})

		t.Run("items", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": "test"}`), &schema)

				if err == nil {
					t.Log("should have error")
					t.FailNow()
				}
			})

			t.Run("should succeed with single schema", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": {"type": "string"}}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should succeed with array of schema", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": [{"type": "string"}]}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should error on bad item", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": ["hi", {"type": "string"}]}`), &schema)

				if err == nil {
					t.Log("should have error")
					t.FailNow()
				}

				if err.Error() != `[items/0] => must be a "Schema"` {
					t.Log(err)
					t.FailNow()
				}
			})
		})

		t.Run("additionalItems", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "additionalItems": "test"}`), &schema)

				if err == nil {
					t.Log("should have error")
					t.FailNow()
				}
			})

			t.Run("should succeed with boolean", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "additionalItems": true}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should succeed with schema", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "additionalItems": {"type": "string"}}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should succeed with nested array schemas", func(t *testing.T) {
				ns := jsonschema.New()
				ns.Read("./schemas/array/nested.json")
				errs := ns.Compile("root")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("minItems", func(t *testing.T) {
			t.Run("should error on string", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "minItems": "test"}`), &schema)

				if err == nil {
					t.Log("should have error")
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "minItems": -3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != "[/minItems] => must be non-negative" {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "minItems": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log("should not have error")
					t.FailNow()
				}
			})
		})

		t.Run("maxItems", func(t *testing.T) {
			t.Run("should error on string", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "maxItems": "test"}`), &schema)

				if err == nil {
					t.Log("should have error")
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "maxItems": -3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != "[/maxItems] => must be non-negative" {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should error when less than minItems", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "minItems": 10, "maxItems": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != `[/maxItems] => must be greater than or equal to "minItems"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "maxItems": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) > 0 {
					t.Log("should not have error")
					t.FailNow()
				}
			})
		})
	})

	t.Run("validate", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should error on string", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/type] => should be type "array"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed on array", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []string{"hello", "world"})

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("items", func(t *testing.T) {
			t.Run("should error on bad item", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": { "type": "string" }}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", 1})

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/1/type] => should be type "string"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should error on extra items", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": [{ "type": "string" }]}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", "world"})

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/1/additionalItems] => undefined array index` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should error on extra items bad type", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": [{ "type": "string" }], "additionalItems": {"type": "number"}}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", "world"})

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/1/type] => should be type "number"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed on extra items bad type", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "items": [{ "type": "string" }], "additionalItems": {"type": "number"}}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", 10})

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("minItems", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "minItems": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", 10})

				if len(errs) != 1 {
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != `[/minItems] => "2" is not greater than or equal to "3"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "minItems": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", 10, true})

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("maxItems", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "maxItems": 1}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", 10})

				if len(errs) != 1 {
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != `[/maxItems] => "2" is not less than or equal to "1"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "maxItems": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", 10, true})

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("uniqueItems", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.ArraySchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "array", "uniqueItems": true}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", []any{"hello", "hello"})

				if len(errs) != 1 {
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != `[/1/uniqueItems] => duplicate item` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})
	})
}
