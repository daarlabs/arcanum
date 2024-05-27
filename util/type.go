package util

import (
	"strconv"
	
	"github.com/daarlabs/arcanum/util/constant/model"
)

func AssertStringToType[T model.Assert](v string) T {
	result := *new(T)
	switch any(result).(type) {
	case string:
		result = any(EscapeString(v)).(T)
	case bool:
		result = any(v == "true").(T)
	case int:
		res, err := strconv.Atoi(v)
		if err == nil {
			result = any(res).(T)
		}
	case float32:
		res, err := strconv.ParseFloat(v, 32)
		if err == nil {
			result = any(float32(res)).(T)
		}
	case float64:
		res, err := strconv.ParseFloat(v, 64)
		if err == nil {
			result = any(res).(T)
		}
	}
	return result
}
