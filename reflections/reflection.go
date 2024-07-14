package reflections

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		fn(val.Field(i).String())
	}
}
