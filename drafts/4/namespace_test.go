package jsonschema_test

import (
	"encoding/json"
	"errors"
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestNamespace(t *testing.T) {
	t.Run("format", func(t *testing.T) {
		ns := jsonschema.New()

		if ns.HasFormat("test") {
			t.FailNow()
		}

		ns.AddFormat("test", func(input string) error {
			if input != "hi" {
				return errors.New("must be hi")
			}

			return nil
		})

		if !ns.HasFormat("test") {
			t.FailNow()
		}

		if err := ns.Format("test", "hi"); err != nil {
			t.Log(err)
			t.FailNow()
		}
	})

	t.Run("schema", func(t *testing.T) {
		ns := jsonschema.New()

		if ns.HasSchema("test") {
			t.FailNow()
		}

		schema := jsonschema.StringSchema{}
		err := json.Unmarshal([]byte(`{"$id": "test", "type": "string", "title": "hello world"}`), &schema)

		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		ns.AddSchema(schema)

		if !ns.HasSchema("test") {
			t.FailNow()
		}

		s := ns.GetSchema("test")

		if s == nil {
			t.Log(s)
			t.FailNow()
		}
	})
}
