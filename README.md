# jsonschema

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/aacebo/jsonschema.svg)](https://pkg.go.dev/github.com/aacebo/jsonschema)
[![Go Report Card](https://goreportcard.com/badge/github.com/aacebo/jsonschema)](https://goreportcard.com/report/github.com/aacebo/jsonschema)
[![Build Status](https://github.com/aacebo/jsonschema/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/aacebo/jsonschema/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/aacebo/jsonschema/graph/badge.svg?token=ZFJMM1BZVM)](https://codecov.io/gh/aacebo/jsonschema)

a zero dependency jsonschema implementation

## Usage

### Files

```go
schema, err := jsonschema.Read("./schema.json")

if err != nil {
    panic(err)
}

errs := jsonschema.Compile(schema)

if len(errs) > 0 {
    panic(errs)
}
```

### Builder

```go
schema := jsonschema.Builder().
    Object().
    Properties(map[string]jsonschema.Schema{
        "test": jsonschema.Builder().String().Build(),
    }).
    AdditionalProperties(jsonschema.Builder().Integer().Build()).
    Required("test").
    Build()

errs := jsonschema.Compile(schema)

if len(errs) > 0 {
    panic(errs)
}

errs = jsonschema.Validate(schema, struct {
    Test  string `json:"test"`
    Other int    `json:"other"`
}{"test", 1})

if len(errs) > 0 {
    panic(errs)
}
```

### Non-Global Namespace

```go
namespace := jsonschema.New()
```

### Custom Keywords

```go
jsonschema.AddKeyword("alphaNum", jsonschema.Keyword{ ... })
```

### Custom Formats

```go
jsonschema.AddFormat("lowercase", func(input string) error {
    if strings.ToLower(input) == input {
        return errors.New("must be lowercase")
    }

    return nil
})
```

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
