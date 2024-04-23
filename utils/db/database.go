package dbutils

import "reflect"

func StructToMap(item interface{}) map[string]interface{} {
	itemValue := reflect.ValueOf(item)
	itemMap := make(map[string]interface{})

	for i := 0; i < itemValue.NumField(); i++ {
		field := itemValue.Type().Field(i)
		tagName := field.Tag.Get("db")
		fieldValue := itemValue.Field(i).Interface()

		if tagName != "" {
			itemMap[tagName] = fieldValue
		}
	}

	return itemMap
}
