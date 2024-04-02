package coerce

import "reflect"

var INT_TYPE = reflect.TypeOf(0)

func Int(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.IsValid() && !value.CanInt() && value.CanConvert(INT_TYPE) {
		value = value.Convert(INT_TYPE)
	}

	return value
}
