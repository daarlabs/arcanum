package form

import (
	"reflect"
)

const (
	valueFieldName = "Value"
)

func CreateStruct[S, R any](src *S) R {
	result := new(R)
	srcRef := reflect.ValueOf(src)
	resultRef := reflect.ValueOf(result)
	for i := 0; i < resultRef.Elem().NumField(); i++ {
		resultField := resultRef.Elem().Field(i)
		resultFieldName := resultRef.Elem().Type().Field(i).Name
		srcField := srcRef.Elem().FieldByName(resultFieldName)
		if srcField.Kind() != reflect.Struct {
			continue
		}
		valueField := srcField.FieldByName(valueFieldName)
		if !valueField.IsValid() {
			continue
		}
		resultField.Set(valueField)
	}
	return *result
}
