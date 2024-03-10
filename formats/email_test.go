package formats_test

import (
	"jsonschema/formats"
	"testing"
)

func TestEmail(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.Email("test")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.Email("test@test.com")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
