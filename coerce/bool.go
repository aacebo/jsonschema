package coerce

import "reflect"

var BOOL_TYPE = reflect.TypeOf(true)

func Bool(value reflect.Value) reflect.Value {
	if value.IsValid() && value.Kind() != reflect.Bool && value.CanConvert(BOOL_TYPE) {
		value = value.Convert(BOOL_TYPE)
	}

	return value
}
