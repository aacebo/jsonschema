package coerce

import "reflect"

var BOOL_TYPE = reflect.TypeOf(true)

func Bool(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.IsValid() && value.Kind() != reflect.Bool && value.CanConvert(BOOL_TYPE) {
		value = value.Convert(BOOL_TYPE)
	}

	return value
}
