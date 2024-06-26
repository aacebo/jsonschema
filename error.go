package jsonschema

import "fmt"

type SchemaError struct {
	Path    string `json:"path"`
	Keyword string `json:"keyword"`
	Message string `json:"message"`
}

func (self SchemaError) Error() string {
	key := self.Path

	if self.Keyword != "" {
		key = fmt.Sprintf("%s/%s", key, self.Keyword)
	}

	return fmt.Sprintf(
		"[%s] => %s",
		key,
		self.Message,
	)
}
