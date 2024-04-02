package coerce

import "reflect"

var FLOAT_TYPE = reflect.TypeOf(0.0)

func Float(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.IsValid() && !value.CanFloat() && value.CanConvert(FLOAT_TYPE) {
		value = value.Convert(FLOAT_TYPE)
	}

	return value
}
