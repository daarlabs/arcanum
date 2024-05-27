package util

import (
	"reflect"
	"strconv"
)

func ConvertSlice(src []string, t interface{}) error {
	v := reflect.ValueOf(t)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Slice {
		return ErrorPointerTarget
	}
	sliceElemType := v.Elem().Type().Elem()
	targetSlice := reflect.MakeSlice(reflect.SliceOf(sliceElemType), 0, len(src))
	for _, item := range src {
		var convertedValue reflect.Value
		switch sliceElemType.Kind() {
		case reflect.Int:
			intVal, err := strconv.Atoi(item)
			if err != nil {
				return err
			}
			convertedValue = reflect.ValueOf(intVal)
		case reflect.Float32:
			floatVal, err := strconv.ParseFloat(item, 32)
			if err != nil {
				return err
			}
			convertedValue = reflect.ValueOf(floatVal)
		case reflect.Float64:
			floatVal, err := strconv.ParseFloat(item, 64)
			if err != nil {
				return err
			}
			convertedValue = reflect.ValueOf(floatVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(item)
			if err != nil {
				return err
			}
			convertedValue = reflect.ValueOf(boolVal)
		case reflect.String:
			convertedValue = reflect.ValueOf(item)
		default:
			return ErrorUnsupportedType
		}
		targetSlice = reflect.Append(targetSlice, convertedValue)
	}
	v.Elem().Set(targetSlice)
	return nil
}

func MustConvertSlice(src []string, t interface{}) {
	if err := ConvertSlice(src, t); err != nil {
		panic(err)
	}
}
