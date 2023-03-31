package converter

import "reflect"

func StructToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	val := reflect.ValueOf(s)
	for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			m[field.Name] = val.Field(i).Interface()
	}
	return m
}
