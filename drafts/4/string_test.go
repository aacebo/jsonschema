package jsonschema_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"jsonschema/core"
	jsonschema "jsonschema/drafts/4"
	"regexp"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("$id", func(t *testing.T) {
		schema := jsonschema.StringSchema{}
		err := json.Unmarshal([]byte(`{"$id": "123", "type": "string"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.GetID() != "123" {
			t.FailNow()
		}
	})

	t.Run("type", func(t *testing.T) {
		schema := jsonschema.StringSchema{}
		err := json.Unmarshal([]byte(`{"type": "string"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.GetType() != core.SCHEMA_TYPE_STRING {
			t.Logf(`expected "string", received "%s"`, schema.GetType())
			t.FailNow()
		}

		if schema.GetID() != "" {
			t.Logf(`expected "", received "%s"`, schema.GetID())
			t.FailNow()
		}
	})

	t.Run("string", func(t *testing.T) {
		schema := jsonschema.StringSchema{}
		err := json.Unmarshal([]byte(`{"type": "string", "title": "hello world"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if schema.String() != `{"type":"string","title":"hello world"}` {
			t.Log(schema)
			t.FailNow()
		}
	})

	t.Run("compile", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string"}`), &schema)

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
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "object", "title": "hello world"}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/type] => "type" must be "string"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})
		})

		t.Run("should succeed with title", func(t *testing.T) {
			schema := jsonschema.StringSchema{}
			err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "title": "hello world"}`), &schema)

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

		t.Run("pattern", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "pattern": "["}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/pattern] => "[" is not a valid regex pattern` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "pattern": "^[0-9]*$"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})
		})

		t.Run("format", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "format": "test"}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/format] => format "test" not found` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "format": "date-time"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("minLength", func(t *testing.T) {
			t.Run("should error on bad type", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": ""}`), &schema)

				if err == nil {
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": -1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/minLength] => "minLength" must be non-negative` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": 0}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})
		})

		t.Run("maxLength", func(t *testing.T) {
			t.Run("should error on bad type", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "maxLength": ""}`), &schema)

				if err == nil {
					t.FailNow()
				}
			})

			t.Run("should error on negative", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "maxLength": -1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/maxLength] => "maxLength" must be non-negative` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should error on lt minLength", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": 3, "maxLength": 1}`), &schema)

				if err != nil {
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Compile("test")

				if len(errs) != 1 {
					t.Logf("should have 1 error, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/maxLength] => "maxLength" must be greater than or equal to "minLength"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "maxLength": 0}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})

			t.Run("should succeed with minLength", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": 1, "maxLength": 3}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}
			})
		})

		t.Run("should error on type and pattern", func(t *testing.T) {
			schema := jsonschema.StringSchema{}
			err := json.Unmarshal([]byte(`{"$id": "test", "type": "object", "pattern": "["}`), &schema)

			if err != nil {
				t.FailNow()
			}

			errs := jsonschema.New().AddSchema(schema).Compile("test")

			if len(errs) != 2 {
				t.Logf("should have 2 errors, received %d", len(errs))
				t.FailNow()
			}

			if errs[0].Error() != `[/type] => "type" must be "string"` {
				t.Log(errs[0])
				t.FailNow()
			}

			if errs[1].Error() != `[/pattern] => "[" is not a valid regex pattern` {
				t.Log(errs[1])
				t.FailNow()
			}
		})
	})

	t.Run("validate", func(t *testing.T) {
		t.Run("type", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", 12)

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/type] => should be type "string"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("pattern", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "pattern": "^[0-9]*$"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/pattern] => "test" does not match pattern "^[0-9]*$"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "pattern": "^[0-9]*$"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "123")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("format", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "format": "test"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/format] => format "test" does not exist` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "format": "date-time"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "2024-03-09T20:39:33.335Z")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should error on custom", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "format": "alphanumeric"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddFormat("alphanumeric", func(input string) error {
					ok := regexp.MustCompile("^[0-9a-zA-Z]*$").MatchString(input)

					if !ok {
						return errors.New(fmt.Sprintf(
							`"%s" is not alphanumeric`,
							input,
						))
					}

					return nil
				}).AddSchema(schema).Validate("test", "!")

				if len(errs) == 0 {
					t.Log("should have error")
					t.FailNow()
				}

				if errs[0].Error() != `[/format] => "!" is not alphanumeric` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed with custom", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "format": "alphanumeric"}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddFormat("alphanumeric", func(input string) error {
					ok := regexp.MustCompile("^[0-9a-zA-Z]*$").MatchString(input)

					if !ok {
						return errors.New(fmt.Sprintf(
							`"%s" is not alphanumeric`,
							input,
						))
					}

					return nil
				}).AddSchema(schema).Validate("test", "abc123")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("minLength", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": 5}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/minLength] => length of "4" is less than min length "5"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": 5}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "tester")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})

		t.Run("maxLength", func(t *testing.T) {
			t.Run("should error", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "maxLength": 5}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "tester")

				if len(errs) != 1 {
					t.Logf("should have 1 errors, received %d", len(errs))
					t.FailNow()
				}

				if errs[0].Error() != `[/maxLength] => length of "6" is greater than max length "5"` {
					t.Log(errs[0])
					t.FailNow()
				}
			})

			t.Run("should succeed", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "maxLength": 5}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})

			t.Run("should succeed with minLength", func(t *testing.T) {
				schema := jsonschema.StringSchema{}
				err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "minLength": 3, "maxLength": 5}`), &schema)

				if err != nil {
					t.Log(err)
					t.FailNow()
				}

				errs := jsonschema.New().AddSchema(schema).Validate("test", "test")

				if len(errs) > 0 {
					t.Log(errs)
					t.FailNow()
				}
			})
		})
	})
}
