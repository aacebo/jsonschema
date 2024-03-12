package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestStringBuilder(t *testing.T) {
	t.Run("should build schema", func(t *testing.T) {
		jsonschema.String().
			ID("root").
			Title("my test title").
			Description("my test desc").
			Pattern("^[0-9]*$").
			Format("date-time").
			MinLength(0).
			MaxLength(100).
			Build()
	})

	t.Run("should build schema in namespace", func(t *testing.T) {
		jsonschema.String().
			ID("root").
			Title("my test title").
			Description("my test desc").
			Pattern("^[0-9]*$").
			Format("date-time").
			MinLength(0).
			MaxLength(100).
			BuildIn(jsonschema.New())
	})
}
