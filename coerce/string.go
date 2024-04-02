package coerce

import "reflect"

var STRING_TYPE = reflect.TypeOf("")

func String(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.IsValid() && value.Kind() != reflect.String && value.CanConvert(STRING_TYPE) {
		value = value.Convert(STRING_TYPE)
	}

	return value
}
