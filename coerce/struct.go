package coerce

import (
	"reflect"
	"strings"
)

func StructFieldName(field reflect.StructField) string {
	if tag := field.Tag.Get("json"); tag != "" {
		parts := strings.Split(tag, ",")

		if len(parts) > 0 {
			if parts[0] == "-" {
				return ""
			} else {
				return parts[0]
			}
		}
	}

	return field.Name
}

func StructFieldByName(object reflect.Value, name string) reflect.Value {
	if object.Kind() != reflect.Struct {
		return reflect.Zero(MAP_TYPE)
	}

	for i := 0; i < object.NumField(); i++ {
		if name == StructFieldName(object.Type().Field(i)) {
			return object.Field(i)
		}
	}

	return reflect.Zero(MAP_TYPE)
}
