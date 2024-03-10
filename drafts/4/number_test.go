package jsonschema_test

import (
	"encoding/json"
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestNumber(t *testing.T) {
	t.Run("$id", func(t *testing.T) {
		schema := jsonschema.NumberSchema{}
		err := json.Unmarshal([]byte(`{"$id": "123", "type": "number"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.GetID() != "123" {
			t.FailNow()
		}
	})

	t.Run("type", func(t *testing.T) {
		schema := jsonschema.NumberSchema{}
		err := json.Unmarshal([]byte(`{"type": "number"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.GetType() != jsonschema.SCHEMA_TYPE_NUMBER {
			t.Logf(`expected "number", received "%s"`, schema.GetType())
			t.FailNow()
		}

		if schema.GetID() != "" {
			t.Logf(`expected "", received "%s"`, schema.GetID())
			t.FailNow()
		}
	})

	t.Run("string", func(t *testing.T) {
		schema := jsonschema.NumberSchema{}
		err := json.Unmarshal([]byte(`{"type": "number", "title": "hello world"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.String() != `{"type":"number","title":"hello world"}` {
			t.Log(schema)
			t.FailNow()
		}
	})

	t.Run("compile", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number"}`), &schema)

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
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "title": "hello world"}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/type] => "type" must be "number"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})

		t.Run("should succeed with title", func(t *testing.T) {
			schema := jsonschema.NumberSchema{}
			err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "title": "hello world"}`), &schema)

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

		t.Run("multipleOf", func(t *testing.T) {
			t.Run("should error on bad type", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "multipleOf": ""}`), &schema)

				if err == nil {
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "multipleOf": -1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != "[/multipleOf] => must be non-negative" {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "multipleOf": 5}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})
		})

		t.Run("minimum", func(t *testing.T) {
			t.Run("should error on bad type", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": ""}`), &schema)

				if err == nil {
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": -1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != "[/minimum] => must be non-negative" {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": 0}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})
		})

		t.Run("maximum", func(t *testing.T) {
			t.Run("should error on bad type", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "maximum": ""}`), &schema)

				if err == nil {
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "maximum": -1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.Log(errs)
					t.FailNow()
				}

				if errs[0].Error() != "[/maximum] => must be non-negative" {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "maximum": 0}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})

			t.Run("should error on lt minimum", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": 3, "maximum": 1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/maximum] => must be greater than or equal to "minimum"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})
	})

	t.Run("validate", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should error on string", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/type] => should be type "number"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed on int", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 12)

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should succeed on float64", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", float64(12.5))

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should succeed on float32", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", float32(12.5))

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("multipleOf", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "multipleOf": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 2)

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/multipleOf] => "2" is not a multiple of "3"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "multipleOf": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 6)

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("minimum", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 2)

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/minimum] => "2" is not greater than or equal to "3"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 3)

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should error on exclusive", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "minimum": 3, "exclusiveMinimum": true}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 3)

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/minimum] => "3" is not greater than or equal to "4"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})

		t.Run("maximum", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "maximum": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 4)

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/maximum] => "4" is not less than or equal to "3"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "maximum": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 3)

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should error on exclusive", func(t *testing.T) {
				schema := jsonschema.NumberSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "number", "maximum": 3, "exclusiveMaximum": true}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 3)

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/maximum] => "3" is not less than or equal to "2"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})
	})
}
