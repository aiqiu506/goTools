package grqStruct

import (
	"reflect"
)

/**
结构体转 map
*/
func StructToMap(i interface{}) map[string]interface{} {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)
	ele := make(map[string]interface{})
	if t.Kind() == reflect.Struct {
		for i, num := 0, v.NumField(); i < num; i++ {
			field := t.Field(i).Name
			if  v.Field(i).Kind()==reflect.Struct{
				ele[field]=StructToMap(v.Field(i).Interface())
			}else{
				ele[field] =v.Field(i).Interface()
			}
		}
	}
	return ele
}
