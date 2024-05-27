package mirage

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/daarlabs/arcanum/config"
)

type Creampuff interface {
	Router
	ErrorHandler(handler Handler) Creampuff
	Layout() Layout
	Run(address string)
	Mux() *http.ServeMux
}

type core struct {
	*router
	*assets
	config       config.Config
	errorHandler Handler
	layout       *layout
	mux          *http.ServeMux
	routes       []*Route
}

const (
	logo = `
░▒▓██████████████▓▒░░▒▓█▓▒░▒▓███████▓▒░ ░▒▓██████▓▒░ ░▒▓██████▓▒░░▒▓████████▓▒░
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓███████▓▒░░▒▓████████▓▒░▒▓█▓▒▒▓███▓▒░▒▓██████▓▒░
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░
░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓████████▓▒░`
)

func New(cfg config.Config) Creampuff {
	mux := http.NewServeMux()
	rts := make([]*Route, 0)
	c := &core{
		config:       cfg,
		errorHandler: defaultErrorHandler,
		layout:       createLayout(),
		mux:          mux,
		routes:       rts,
	}
	c.router = &router{
		config: cfg,
		mux:    mux,
		prefix: cfg.Router.Prefix,
		routes: &rts,
	}
	c.assets = &assets{
		dir:    cfg.App.Assets,
		public: cfg.App.Public,
	}
	c.router.core = c
	c.router.createGetWildcardRoute()
	c.onInit()
	return c
}

func (c *core) ErrorHandler(handler Handler) Creampuff {
	c.errorHandler = handler
	return c
}

func (c *core) Layout() Layout {
	return c.layout
}

func (c *core) Run(address string) {
	fmt.Println(logo)
	fmt.Println("")
	fmt.Println("Name: ", c.config.App.Name)
	fmt.Println("Version: ", "0.1.0")
	log.Fatalln(http.ListenAndServe(address, c.mux))
}

func (c *core) Mux() *http.ServeMux {
	return c.mux
}

func (c *core) onInit() {
	go func() {
		time.Sleep(2 * time.Second)
		c.assets.mustRead()
	}()
}
