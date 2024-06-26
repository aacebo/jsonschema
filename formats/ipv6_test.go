package formats_test

import (
	"testing"

	"github.com/aacebo/jsonschema/formats"
)

func TestIPv6(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.IPv6("test")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should error on ipv4", func(t *testing.T) {
		err := formats.IPv6("127.0.0.1")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.IPv6("2603:7000:873a:2a48:3d6c:bce7:fe8d:3b8")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
