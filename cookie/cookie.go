package cookie

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/daarlabs/arcanum/env"
)

type Cookie interface {
	Get(name string) string
	Set(name string, value any, expiration time.Duration)
	Destroy(name string)
}

type cookie struct {
	req  *http.Request
	res  http.ResponseWriter
	path string
}

func New(
	req *http.Request,
	res http.ResponseWriter,
	path string,
) Cookie {
	return &cookie{
		req:  req,
		res:  res,
		path: path,
	}
}

func (c cookie) Get(name string) string {
	r, err := c.req.Cookie(name)
	if err != nil {
		return ""
	}
	return r.Value
}

func (c cookie) Set(name string, value any, expiration time.Duration) {
	// domain := c.req.Header.Get("Origin")
	// if strings.Contains(domain, "//") {
	// 	domain = domain[strings.Index(domain, "//")+2:]
	// }
	http.SetCookie(
		c.res, &http.Cookie{
			Name: name,
			// Domain:  domain,
			Value:   fmt.Sprintf("%v", value),
			Path:    c.path,
			Expires: time.Now().Add(expiration),
			Secure:  env.Production(),
		},
	)
}

func (c cookie) Destroy(name string) {
	c.Set(name, "", time.Millisecond)
}
