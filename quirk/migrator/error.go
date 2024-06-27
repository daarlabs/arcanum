package migrator

import "errors"

var (
	ErrorInvalidDatabase = errors.New("invalid database")
	ErrorInvalidDir      = errors.New("invalid directory")
)
