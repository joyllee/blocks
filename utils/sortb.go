package utils

import "reflect"

func SortCompare(a interface{}, b interface{}, key string, sortStr string) bool {
	vi := reflect.ValueOf(a).FieldByName(key).Interface()
	vj := reflect.ValueOf(b).FieldByName(key).Interface()
	if sortStr == "asc" {
		switch reflect.ValueOf(a).FieldByName(key).Kind() {
		case reflect.String:
			return vi.(string) < vj.(string)
		case reflect.Int:
			return vi.(int) < vj.(int)
		case reflect.Int64:
			return vi.(int64) < vj.(int64)
		case reflect.Float64:
			return vi.(float64) < vj.(float64)
		case reflect.Float32:
			return vi.(float32) < vj.(float32)
		}
	} else {
		switch reflect.ValueOf(a).FieldByName(key).Kind() {
		case reflect.String:
			return vi.(string) > vj.(string)
		case reflect.Int:
			return vi.(int) > vj.(int)
		case reflect.Int64:
			return vi.(int64) > vj.(int64)
		case reflect.Float64:
			return vi.(float64) > vj.(float64)
		case reflect.Float32:
			return vi.(float32) > vj.(float32)
		}
	}
	return false
}
