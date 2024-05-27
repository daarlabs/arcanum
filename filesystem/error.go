package filesystem

import "errors"

var (
	ErrorMissingCloud = errors.New("cloud client is not initialized")
)
