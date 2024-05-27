package util

import (
	"reflect"
	"strconv"
)

func ConvertValue(src string, t interface{}) error {
	tv := reflect.ValueOf(t)
	if tv.Kind() != reflect.Ptr {
		return ErrorPointerTarget
	}
	switch reflect.TypeOf(t).Elem().Kind() {
	case reflect.Int:
		val, err := strconv.Atoi(src)
		if err != nil {
			return err
		}
		tv.Elem().Set(reflect.ValueOf(val))
	case reflect.Float32:
		val, err := strconv.ParseFloat(src, 32)
		if err != nil {
			return err
		}
		tv.Elem().Set(reflect.ValueOf(val))
	case reflect.Float64:
		val, err := strconv.ParseFloat(src, 64)
		if err != nil {
			return err
		}
		tv.Elem().Set(reflect.ValueOf(val))
	case reflect.Bool:
		val, err := strconv.ParseBool(src)
		if err != nil {
			return err
		}
		tv.Elem().Set(reflect.ValueOf(val))
	case reflect.String:
		tv.Elem().Set(reflect.ValueOf(src))
	default:
		return ErrorUnsupportedType
	}
	return nil
}

func MustConvertValue(src string, t interface{}) {
	if err := ConvertValue(src, t); err != nil {
		panic(err)
	}
}
