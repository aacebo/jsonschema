# jsonschema

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/aacebo/jsonschema?status.svg)](https://pkg.go.dev/github.com/aacebo/jsonschema)
[![Go Report Card](https://goreportcard.com/badge/github.com/aacebo/jsonschema)](https://goreportcard.com/report/github.com/aacebo/jsonschema)
[![Build Status](https://github.com/aacebo/jsonschema/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/aacebo/jsonschema/actions/workflows/ci.yml)

a zero dependency jsonschema implementation

## CLI

```bash
$: jsonschema **/*
```

## Supported Drafts

| Draft   | Status  |
|---------|---------|
| 4       | ✅      |
| 6       | ✅      |
| 7       | ✅      |
| 2019-09 | ⏳      |
| 2020-12 | ⏳      |

## To Do

- schema builder pattern
- keywords:
    - object
        - properties
        - propertyNames
        - patternProperties
        - additionalProperties
        - unevaluatedProperties
        - minProperties
        - maxProperties
        - required
    - array
        - prefixItems
        - unevaluatedItems
        - minContains
        - maxContains
    - deprecated
    - readOnly
    - writeOnly
    - examples
    - contentMediaType
    - contentEncoding
    - if/then/else
