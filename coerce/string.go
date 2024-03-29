package coerce

import "reflect"

var STRING_TYPE = reflect.TypeOf("")

func String(value reflect.Value) reflect.Value {
	if value.IsValid() && value.Kind() != reflect.String && value.CanConvert(STRING_TYPE) {
		value = value.Convert(STRING_TYPE)
	}

	return value
}
