package mirage

import (
	"errors"
)

var (
	ErrorInvalidDatabase = errors.New("invalid database")
	ErrorNoPtr           = errors.New("target is not a pointer")
	ErrorMismatchType    = errors.New("target is not equal with value type")
)
