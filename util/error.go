package util

import "errors"

var (
	ErrorPointerTarget   = errors.New("target must be a pointer")
	ErrorUnsupportedType = errors.New("unsupported type")
)
