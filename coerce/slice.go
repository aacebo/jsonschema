package coerce

import "reflect"

var SLICE_TYPE = reflect.TypeOf([]any{})

func Slice(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.IsValid() && value.Kind() != reflect.Slice && value.CanConvert(SLICE_TYPE) {
		value = value.Convert(SLICE_TYPE)
	}

	return value
}

func SliceOf(value reflect.Value, t reflect.Type) reflect.Value {
	_type := reflect.SliceOf(t)

	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.IsValid() && value.Type() != _type && value.CanConvert(_type) {
		value = value.Convert(_type)
	}

	return value
}
