package csrf

import (
	"net/http"
	"strings"
	"time"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type Csrf interface {
	Exists(name, value string) (bool, error)
	Get(name, value string) (Token, error)
	Create(token Token) (string, error)
	Destroy(token Token) error
	Clean(ignore string) error
	
	MustExists(name, value string) bool
	MustGet(name, value string) Token
	MustCreate(token Token) string
	MustDestroy(token Token)
	MustClean(ignore string)
}

type csrf struct {
	Cache      cache.Client
	Cookie     cookie.Cookie
	Enabled    bool
	Expiration time.Duration
	Request    *http.Request
}

type Token struct {
	Exists    bool   `json:"exists"`
	Name      string `json:"name"`
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	Value     string `json:"value"`
}

const (
	cookieKey = "X-Csrf"
)

var (
	defaultExpiration = time.Hour
)

func New(configs ...Config) Csrf {
	r := &csrf{
		Enabled:    true,
		Expiration: defaultExpiration,
	}
	for _, item := range configs {
		c, ok := item.(*config)
		if !ok {
			continue
		}
		switch c.name {
		case configCache:
			r.Cache = c.value.(cache.Client)
		case configCookie:
			r.Cookie = c.value.(cookie.Cookie)
		case configEnabled:
			r.Enabled = c.value.(bool)
		case configExpiration:
			r.Expiration = c.value.(time.Duration)
		case configRequest:
			r.Request = c.value.(*http.Request)
		}
	}
	return r
}

func (c *csrf) Exists(name, value string) (bool, error) {
	ok := c.Cache.Exists(c.createCacheKey(Token{Name: name, Value: value}))
	if !ok {
		return ok, ErrorInvalidToken
	}
	return ok, nil
}

func (c *csrf) MustExists(name, value string) bool {
	r, err := c.Exists(name, value)
	if err != nil {
		panic(err)
	}
	return r
}

func (c *csrf) Get(name, value string) (Token, error) {
	var result Token
	if err := c.Cache.Get(c.createCacheKey(Token{Name: name, Value: value}), &result); err != nil {
		return result, err
	}
	if !result.Exists {
		return result, ErrorInvalidToken
	}
	return result, nil
}

func (c *csrf) MustGet(name, value string) Token {
	r, err := c.Get(name, value)
	if err != nil {
		panic(err)
	}
	return r
}

func (c *csrf) Create(token Token) (string, error) {
	token.Exists = true
	token.Value = uniuri.New()
	c.Cookie.Set(cookieKey+"-"+token.Name, token.Value, c.Expiration)
	if err := c.Cache.Set(
		c.createCacheKey(token),
		token,
		c.Expiration,
	); err != nil {
		return "", err
	}
	return token.Value, nil
}

func (c *csrf) MustCreate(token Token) string {
	r, err := c.Create(token)
	if err != nil {
		panic(err)
	}
	return r
}

func (c *csrf) Destroy(token Token) error {
	return c.Cache.Destroy(c.createCacheKey(token))
}

func (c *csrf) Clean(ignore string) error {
	cookies := c.Request.Header.Get("Cookie")
	if cookies == "" {
		return ErrorMissingCookies
	}
	for _, part := range strings.Split(cookies, ";") {
		part = strings.TrimSpace(part)
		if !strings.Contains(part, cookieKey) {
			continue
		}
		if !strings.Contains(part, "=") {
			continue
		}
		if len(part) < strings.Index(part, "=")+1 {
			continue
		}
		name := part[:strings.Index(part, "=")]
		token := part[strings.Index(part, "=")+1:]
		if err := c.Destroy(Token{Name: strings.TrimPrefix(name, cookieKey+"-"), Value: token}); err != nil {
			return err
		}
		if ignore != strings.TrimPrefix(name, cookieKey+"-") {
			c.Cookie.Destroy(name)
		}
	}
	return nil
}

func (c *csrf) MustDestroy(token Token) {
	err := c.Destroy(token)
	if err != nil {
		panic(err)
	}
}

func (c *csrf) MustClean(ignore string) {
	err := c.Clean(ignore)
	if err != nil {
		panic(err)
	}
}

func (c *csrf) createCacheKey(token Token) string {
	if len(token.Name) > 0 {
		return "csrf:" + token.Name + ":" + token.Value
	}
	return "csrf:" + token.Value
}
