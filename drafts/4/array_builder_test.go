package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestArrayBuilder(t *testing.T) {
	t.Run("should build schema", func(t *testing.T) {
		jsonschema.Array().
			ID("root").
			Title("my test title").
			Description("my test desc").
			Items(
				jsonschema.String().Build(),
				jsonschema.Number().Build(),
			).
			AdditionalItems(jsonschema.Null().Build()).
			MinItems(0).
			MaxItems(100).
			UniqueItems(true).
			Build()
	})

	t.Run("should build schema in namespace", func(t *testing.T) {
		jsonschema.Array().
			ID("root").
			Title("my test title").
			Description("my test desc").
			Item(jsonschema.String().Build()).
			AdditionalItemsAllowed(false).
			MinItems(0).
			MaxItems(100).
			UniqueItems(true).
			BuildIn(jsonschema.New())
	})
}
