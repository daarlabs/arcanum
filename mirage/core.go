package mirage

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	
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

const (
	Version = "0.1.0"
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
	c.assets = createAssets(cfg)
	c.router = &router{
		config: cfg,
		mux:    mux,
		prefix: cfg.Router.Prefix,
		routes: &rts,
		assets: c.assets,
	}
	c.router.core = c
	c.router.createWildcardRoute()
	c.router.createDynamicAssetsRoute()
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
	if strings.HasPrefix(address, ":") {
		address = "localhost" + address
	}
	fmt.Println(logo)
	fmt.Println("")
	fmt.Println("Name: ", c.config.App.Name)
	fmt.Println("Address: ", address)
	fmt.Println("Version: ", Version)
	log.Fatalln(http.ListenAndServe(address, c.mux))
}

func (c *core) Mux() *http.ServeMux {
	return c.mux
}

func (c *core) onInit() {
	c.assets.mustProcess()
}
