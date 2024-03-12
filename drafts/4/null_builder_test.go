package jsonschema_test

import (
	jsonschema "jsonschema/drafts/4"
	"testing"
)

func TestNullBuilder(t *testing.T) {
	t.Run("should build schema", func(t *testing.T) {
		jsonschema.Null().
			ID("root").
			Title("my test title").
			Description("my test desc").
			Build()
	})

	t.Run("should build schema in namespace", func(t *testing.T) {
		jsonschema.Null().
			ID("root").
			Title("my test title").
			Description("my test desc").
			BuildIn(jsonschema.New())
	})
}
