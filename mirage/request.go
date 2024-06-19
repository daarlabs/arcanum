package mirage

import (
	"net/http"
	"net/url"
	
	"github.com/daarlabs/arcanum/hx"
	
	"github.com/daarlabs/arcanum/util/constant/header"
)

type Request interface {
	Action() string
	ContentType() string
	Form() url.Values
	Header() http.Header
	Host() string
	Ip() string
	Is() RequestIs
	Method() string
	Name() string
	Origin() string
	Parsed() Map
	Path() string
	PathValue(key string, defaultValue ...string) string
	QueryParam(key string, defaultValue ...string) string
	Protocol() string
	Raw() *http.Request
	UserAgent() string
}

type RequestIs interface {
	Get() bool
	Post() bool
	Put() bool
	Patch() bool
	Delete() bool
	Action(actionName ...string) bool
	Active(routeName string) bool
	Hx() bool
	Options() bool
	Head() bool
	Connect() bool
	Trace() bool
}

type request struct {
	r            *http.Request
	componentCtx *componentCtx
	route        *Route
	parsed       Map
}

type requestIs struct {
	r            *http.Request
	componentCtx *componentCtx
	route        *Route
}

func (r request) Action() string {
	return r.r.URL.Query().Get(Action)
}

func (r request) ContentType() string {
	return r.r.Header.Get(header.ContentType)
}

func (r request) Form() url.Values {
	return r.r.Form
}

func (r request) Header() http.Header {
	return r.r.Header
}

func (r request) Host() string {
	return r.Protocol() + "://" + r.r.Host
}

func (r request) Ip() string {
	return r.r.Header.Get("X-Forwarded-For")
}

func (r request) Is() RequestIs {
	return requestIs{
		r:            r.r,
		componentCtx: r.componentCtx,
		route:        r.route,
	}
}

func (r request) Method() string {
	return r.r.Method
}

func (r request) Name() string {
	if r.route == nil {
		return ""
	}
	return r.route.Name
}

func (r request) Origin() string {
	return r.r.Header.Get(header.Origin)
}

func (r request) Parsed() Map {
	return r.parsed
}

func (r request) Path() string {
	return r.r.URL.Path
}

func (r request) PathValue(key string, defaultValue ...string) string {
	value := r.r.PathValue(key)
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func (r request) QueryParam(key string, defaultValue ...string) string {
	value := r.r.URL.Query().Get(key)
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func (r request) Protocol() string {
	if r.r.TLS == nil {
		return "http"
	}
	return "https"
}

func (r request) Raw() *http.Request {
	return r.r
}

func (r request) UserAgent() string {
	return r.r.Header.Get(header.UserAgent)
}

func (r requestIs) Get() bool {
	return r.r.Method == http.MethodGet
}

func (r requestIs) Post() bool {
	return r.r.Method == http.MethodPost
}

func (r requestIs) Put() bool {
	return r.r.Method == http.MethodPut
}

func (r requestIs) Patch() bool {
	return r.r.Method == http.MethodPatch
}

func (r requestIs) Delete() bool {
	return r.r.Method == http.MethodDelete
}

func (r requestIs) Active(routeName string) bool {
	return r.route.Name == routeName
}

func (r requestIs) Action(actionName ...string) bool {
	action := r.r.URL.Query().Get(Action)
	isAction := len(action) > 0
	if isAction && len(actionName) > 0 && r.componentCtx != nil {
		actionPrefix := r.route.Name + "_" + r.componentCtx.name
		var equal bool
		for _, an := range actionName {
			if actionPrefix+"_"+an == action {
				equal = true
				break
			}
		}
		return equal
	}
	return isAction
}

func (r requestIs) Hx() bool {
	return r.r.Header.Get(hx.RequestHeaderRequest) == "true"
}

func (r requestIs) Options() bool {
	return r.r.Method == http.MethodOptions
}

func (r requestIs) Head() bool {
	return r.r.Method == http.MethodHead
}

func (r requestIs) Connect() bool {
	return r.r.Method == http.MethodConnect
}

func (r requestIs) Trace() bool {
	return r.r.Method == http.MethodTrace
}
