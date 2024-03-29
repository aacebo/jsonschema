package coerce

import "reflect"

var INT_TYPE = reflect.TypeOf(0)

func Int(value reflect.Value) reflect.Value {
	if value.IsValid() && !value.CanInt() && value.CanConvert(INT_TYPE) {
		value = value.Convert(INT_TYPE)
	}

	return value
}
