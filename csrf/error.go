package csrf

import "errors"

var (
	ErrorInvalidToken   = errors.New("invalid csrf token")
	ErrorMissingCookies = errors.New("missing http header cookies")
)
