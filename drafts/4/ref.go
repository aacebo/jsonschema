package jsonschema

import (
	"fmt"
	"jsonschema/core"
	"net/url"
	"strings"
)

// https://json-schema.org/understanding-json-schema/structuring#ref
type RefSchema struct {
	Path string `json:"$ref"`
}

func (self RefSchema) GetID() string {
	return ""
}

func (self RefSchema) GetType() core.SchemaType {
	return core.SCHEMA_TYPE_REF
}

func (self RefSchema) String() string {
	return self.Path
}

func (self RefSchema) Parse(base string) (RefResult, error) {
	if self.Path == "" {
		return RefResult{URL: base}, nil
	}

	refURL, err := url.Parse(self.Path)

	if err != nil {
		return RefResult{}, err
	}

	if refURL.IsAbs() {
		return RefResult{URL: self.Path}, nil
	}

	baseURL, err := url.Parse(base)

	if err != nil {
		return RefResult{}, err
	}

	if strings.HasPrefix(refURL.String(), "#") {
		return RefResult{
			URL:  baseURL.String(),
			Path: refURL.String(),
		}, nil
	}

	return RefResult{
		URL: baseURL.ResolveReference(refURL).String(),
	}, nil
}

func (self RefSchema) compile(ns core.Namespace[Schema], id string, path string) []core.SchemaError {
	ref, err := self.Parse(id)

	if err != nil {
		return []core.SchemaError{{
			Path:    path,
			Keyword: "$ref",
			Message: fmt.Sprintf(
				`failed to parse reference "%s"`,
				ref,
			),
		}}
	}

	schema, err := ns.Resolve(ref.URL, ref.Path)

	if err != nil {
		return []core.SchemaError{{
			Path:    path,
			Keyword: "$ref",
			Message: fmt.Sprintf(
				`failed to resolve reference "%s"`,
				ref,
			),
		}}
	}

	return schema.compile(ns, "", path)
}

func (self RefSchema) validate(ns core.Namespace[Schema], id string, path string, value any) []core.SchemaError {
	ref, err := self.Parse(id)

	if err != nil {
		return []core.SchemaError{{
			Path:    path,
			Keyword: "$ref",
			Message: fmt.Sprintf(
				`failed to parse reference "%s"`,
				ref,
			),
		}}
	}

	schema, err := ns.Resolve(ref.URL, ref.Path)

	if err != nil {
		return []core.SchemaError{{
			Path:    path,
			Keyword: "$ref",
			Message: fmt.Sprintf(
				`failed to resolve reference "%s"`,
				ref,
			),
		}}
	}

	return schema.validate(ns, "", path, value)
}

type RefResult struct {
	URL  string
	Path string
}

func (self RefResult) String() string {
	return self.URL + self.Path
}
