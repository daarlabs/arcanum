package mirage

import (
	"errors"
)

var (
	ErrorInvalidDatabase = errors.New("invalid database")
	ErrorNoPtr           = errors.New("target is not a pointer")
	ErrorMismatchType    = errors.New("target is not equal with value type")
	ErrorInvalidLayout   = errors.New("invalid layout")
)

func defaultErrorHandler(c Ctx) error {
	return c.Response().
		Status(c.Response().Intercept().Status()).
		Error(c.Response().Intercept().Error())
}
