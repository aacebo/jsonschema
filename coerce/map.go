package coerce

import "reflect"

var MAP_TYPE = reflect.TypeOf(map[string]any{})

func Map(value reflect.Value) reflect.Value {
	if value.IsValid() && value.Kind() != reflect.Map && value.CanConvert(MAP_TYPE) {
		value = value.Convert(MAP_TYPE)
	}

	return value
}

func MapOf(value reflect.Value, t reflect.Type) reflect.Value {
	_type := reflect.MapOf(STRING_TYPE, t)

	if value.IsValid() && value.Type() != _type && value.CanConvert(_type) {
		value = value.Convert(_type)
	}

	return value
}