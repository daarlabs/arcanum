package auth

import "fmt"

const (
	SessionCacheKey = "session"
	TfaCacheKey     = "tfa"
)

func createSessionCacheKey(token string) string {
	return fmt.Sprintf("%s:%s", SessionCacheKey, token)
}

func createTfaCacheKey(token string) string {
	return fmt.Sprintf("%s:%s", TfaCacheKey, token)
}
