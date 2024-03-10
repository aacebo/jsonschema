package core

import (
	"net/url"
	"strings"
)

// https://json-schema.org/understanding-json-schema/structuring#ref
func ResolveRef(base string, ref string) (string, error) {
	if ref == "" {
		return base, nil
	}

	if strings.HasPrefix(ref, "urn:") {
		return ref, nil
	}

	refURL, err := url.Parse(ref)

	if err != nil {
		return "", err
	}

	if refURL.IsAbs() {
		return ref, nil
	}

	if strings.HasPrefix(base, "urn:") {
		base, _ = split(base)
		return base + ref, nil
	}

	baseURL, err := url.Parse(base)

	if err != nil {
		return "", err
	}

	return baseURL.ResolveReference(refURL).String(), nil
}

func split(uri string) (string, string) {
	hash := strings.IndexByte(uri, '#')

	if hash == -1 {
		return uri, "#"
	}
	f := uri[hash:]

	if f == "#/" {
		f = "#"
	}

	return uri[0:hash], f
}
