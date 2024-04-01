# jsonschema

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/aacebo/jsonschema.svg)](https://pkg.go.dev/github.com/aacebo/jsonschema)
[![Go Report Card](https://goreportcard.com/badge/github.com/aacebo/jsonschema)](https://goreportcard.com/report/github.com/aacebo/jsonschema)
[![Build Status](https://github.com/aacebo/jsonschema/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/aacebo/jsonschema/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/aacebo/jsonschema/graph/badge.svg?token=ZFJMM1BZVM)](https://codecov.io/gh/aacebo/jsonschema)

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

## Features

| Name                  | Status    |
|-----------------------|-----------|
| Custom Keywords       | ✅        |
| Custom Error Messages | ⏳        |
| Custom Formats        | ✅        |

## To Do

- schema builder pattern
- cli
- struct tags
- add message keyword for custom error messages
- keywords:
    - object
        - unevaluatedProperties
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
