package mirage

import (
	"context"
	"net/http"
	"sync"
	
	"github.com/daarlabs/arcanum/filesystem"
	"github.com/daarlabs/arcanum/tempest"
	
	"github.com/daarlabs/arcanum/auth"
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/config"
	"github.com/daarlabs/arcanum/cookie"
	"github.com/daarlabs/arcanum/csrf"
	"github.com/daarlabs/arcanum/mailer"
	"github.com/daarlabs/arcanum/parser"
	"github.com/daarlabs/arcanum/quirk"
	"github.com/daarlabs/arcanum/sender"
)

type Ctx interface {
	Auth(dbname ...string) auth.Manager
	Cache() cache.Client
	Config() config.Config
	Continue() error
	Cookie() cookie.Cookie
	Create() Factory
	Csrf() csrf.Csrf
	DB(dbname ...string) *quirk.DB
	Email() mailer.Mailer
	Export() Export
	Files() filesystem.Client
	Flash() Flash
	Generate() Generator
	Lang() Lang
	Page() Page
	Parse() parser.Parse
	Request() Request
	Response() Response
	Tempest() tempest.Class
	Translate(key string, args ...map[string]any) string
}

type ctx struct {
	context.Context
	err              error
	cachedComponents *map[string]MandatoryComponent
	config           config.Config
	cookie           cookie.Cookie
	csrf             csrf.Csrf
	files            filesystem.Client
	mu               *sync.Mutex
	page             *page
	r                *http.Request
	w                http.ResponseWriter
	route            *Route
	routes           *[]*Route
	response         *response
	state            *state
	assets           *assets
	lang             *lang
	component        *componentCtx
	tempest          *tempest.Context
	write            *bool
}

type ctxParam struct {
	assets           *assets
	cachedComponents *map[string]MandatoryComponent
	config           config.Config
	layout           *layout
	matchedRoute     *Route
	routes           *[]*Route
	r                *http.Request
	w                http.ResponseWriter
}

func createContext(p ctxParam) *ctx {
	cx := context.Background()
	write := true
	c := &ctx{
		Context:          cx,
		cachedComponents: p.cachedComponents,
		config:           p.config,
		files:            filesystem.New(cx, p.config.Filesystem),
		mu:               &sync.Mutex{},
		page:             createPage(),
		route:            p.matchedRoute,
		routes:           p.routes,
		r:                p.r,
		w:                p.w,
		assets:           p.assets,
		write:            &write,
	}
	if c.config.Tempest != nil {
		c.tempest = c.config.Tempest.Context()
	}
	c.cookie = cookie.New(c.r, c.w, c.createCookiePathBasedOnRouterCookiePrefix())
	c.csrf = csrf.New(
		csrf.Cache(c.Cache()),
		csrf.Cookie(c.cookie),
		csrf.Request(p.r),
	)
	c.lang = createLang(c.Config(), c.Request(), c.Cookie())
	c.response = &response{
		Sender: sender.New(p.r, p.w, &write),
		ctx:    c,
		layout: p.layout,
		l:      p.layout.factories[Main],
	}
	c.state = createState(c.Cache(), c.Cookie())
	return c
}

func (c *ctx) Auth(dbname ...string) auth.Manager {
	var db *quirk.DB
	var ok bool
	dbn := Main
	if len(dbname) > 0 {
		dbn = dbname[0]
	}
	if len(c.config.Database) > 0 {
		db, ok = c.config.Database[dbn]
		if !ok {
			panic(ErrorInvalidDatabase)
		}
	}
	return auth.New(
		db,
		c.r,
		c.w,
		c.cookie,
		c.Cache(),
		c.config.Security.Auth,
	)
}

func (c *ctx) Cache() cache.Client {
	return cache.New(c.Context, c.config.Cache.Memory, c.config.Cache.Redis)
}

func (c *ctx) Config() config.Config {
	return c.config
}

func (c *ctx) Cookie() cookie.Cookie {
	return c.cookie
}

func (c *ctx) Continue() error {
	// c.mu.Unlock()
	return nil
}

func (c *ctx) Create() Factory {
	return factory{ctx: c}
}

func (c *ctx) Csrf() csrf.Csrf {
	return c.csrf
}

func (c *ctx) DB(dbname ...string) *quirk.DB {
	dbn := Main
	if len(dbname) > 0 {
		dbn = dbname[0]
	}
	db, ok := c.config.Database[dbn]
	if !ok {
		panic(ErrorInvalidDatabase)
	}
	return db
}

func (c *ctx) Email() mailer.Mailer {
	return mailer.New(c.config.Smtp)
}

func (c *ctx) Export() Export {
	return createExport(c.config.Export)
}

func (c *ctx) Files() filesystem.Client {
	return c.files
}

func (c *ctx) Flash() Flash {
	return flash{state: c.state}
}

func (c *ctx) Generate() Generator {
	return &generator{c}
}

func (c *ctx) Lang() Lang {
	return c.lang
}

func (c *ctx) Page() Page {
	return c.page
}

func (c *ctx) Parse() parser.Parse {
	return parser.New(c.r, []byte{}, c.config.Parser.Limit)
}

func (c *ctx) Request() Request {
	return request{c.r, c.route}
}

func (c *ctx) Response() Response {
	return c.response
}

func (c *ctx) Translate(key string, args ...map[string]any) string {
	if !c.config.Localization.Enabled {
		return key
	}
	return c.config.Localization.Translator.Translate(c.Lang().Current(), key, args...)
}

func (c *ctx) Tempest() tempest.Class {
	if c.config.Tempest == nil {
		return nil
	}
	return c.tempest.Class()
}

func (c *ctx) createCookiePathBasedOnRouterCookiePrefix() string {
	if len(c.config.Router.Prefix.Cookie) > 0 {
		return c.config.Router.Prefix.Cookie
	}
	return "/"
}
