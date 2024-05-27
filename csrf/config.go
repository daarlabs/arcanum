package csrf

import (
	"net/http"
	"time"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type Config interface{}

type config struct {
	name  string
	value any
}

const (
	configCache      = "cache"
	configCookie     = "cookie"
	configEnabled    = "enabled"
	configExpiration = "name"
	configRequest    = "request"
)

func Cache(client cache.Client) Config {
	return &config{
		name:  configCache,
		value: client,
	}
}

func Cookie(cookie cookie.Cookie) Config {
	return &config{
		name:  configCookie,
		value: cookie,
	}
}

func Enabled(enabled ...bool) Config {
	value := true
	if len(enabled) > 0 {
		value = enabled[0]
	}
	return &config{
		name:  configEnabled,
		value: value,
	}
}

func Expiration(expiration time.Duration) Config {
	return &config{
		name:  configExpiration,
		value: expiration,
	}
}

func Request(request *http.Request) Config {
	return &config{
		name:  configRequest,
		value: request,
	}
}
