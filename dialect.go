package jsonschema

import "reflect"

type Dialect string

const (
	DIALECT_DRAFT_4 Dialect = "http://json-schema.org/draft-04/schema#"
	DIALECT_DRAFT_6 Dialect = "http://json-schema.org/draft-06/schema#"
	DIALECT_DRAFT_7 Dialect = "http://json-schema.org/draft-07/schema#"
)

func (self Dialect) Valid() bool {
	switch self {
	case DIALECT_DRAFT_4:
		fallthrough
	case DIALECT_DRAFT_6:
		fallthrough
	case DIALECT_DRAFT_7:
		return true
	}

	return false
}

// https://json-schema.org/understanding-json-schema/reference/schema#declaring-a-dialect
func dialect(key string) Keyword {
	return Keyword{
		Compile: func(ns *Namespace, ctx Context, config reflect.Value) []SchemaError {
			errs := []SchemaError{}

			if config.Kind() != reflect.String {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a "string"`,
				})

				return errs
			}

			if !Dialect(config.String()).Valid() {
				errs = append(errs, SchemaError{
					Path:    ctx.Path,
					Keyword: key,
					Message: `must be a supported "Dialect"`,
				})
			}

			return errs
		},
	}
}
