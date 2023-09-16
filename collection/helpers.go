package collection

import (
	"fmt"
	"reflect"
)

func getFieldValueFloat32(data interface{}, fieldName string) float64 {
	val := reflect.ValueOf(data)

	// Make sure the data is a struct
	if val.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Error: Data is not a struct\n"))
	}

	// Get the field by name
	field := val.FieldByName(fieldName)

	// Check if the field exists
	if !field.IsValid() {
		panic(fmt.Sprintf("Error: Field '%s' not found in the struct\n", fieldName))
	}

	switch r := field.Interface().(type) {
	case int:
		return float64(r)
	case int32:
		return float64(r)
	case int64:
		return float64(r)
	case float32:
		return float64(r)
	case float64:
		return r
	default:
		panic(fmt.Sprintf("Error: Field '%s' not a number field\n", fieldName))
	}
	return (field.Interface()).(float64)
}
