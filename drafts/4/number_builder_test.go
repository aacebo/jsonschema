package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestNumberBuilder(t *testing.T) {
	t.Run("should build schema", func(t *testing.T) {
		jsonschema.Number().
			ID("root").
			Title("my test title").
			Description("my test desc").
			MultipleOf(10).
			Minimum(0).
			Maximum(100).
			Build()
	})

	t.Run("should build schema in namespace", func(t *testing.T) {
		jsonschema.Number().
			ID("root").
			Title("my test title").
			Description("my test desc").
			MultipleOf(10).
			Minimum(0).
			Maximum(100).
			BuildIn(jsonschema.New())
	})
}
