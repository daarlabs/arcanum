package devtool

import (
	"embed"
	"fmt"
	"log"
	
	"github.com/daarlabs/arcanum/devtool/hub"
	"github.com/daarlabs/arcanum/gox"
)

const (
	HubPath = "/_development"
)

//go:embed style.css
var style embed.FS

//go:embed script.js
var script embed.FS

type Devtool struct {
	script string
	style  string
	hub    *hub.Hub
}

func New() *Devtool {
	d := &Devtool{
		hub: hub.New(),
	}
	d.loadStyle()
	d.loadScript()
	return d
}

func (d *Devtool) Hub() *hub.Hub {
	return d.hub
}

func (d *Devtool) CreateView(
	renderDuration string, values []any, queries []Query, session map[string]any, route map[string]any,
) gox.Node {
	return createView(
		d.style,
		d.script,
		renderDuration,
		values,
		queries,
		session,
		route,
	)
}

func (d *Devtool) loadScript() {
	scriptBytes, err := script.ReadFile("script.js")
	if err != nil {
		log.Fatalln(fmt.Errorf("error while reading script.js: %w", err))
	}
	d.script = string(scriptBytes)
}

func (d *Devtool) loadStyle() {
	styleBytes, err := style.ReadFile("style.css")
	if err != nil {
		log.Fatalln(fmt.Errorf("error while reading style.css: %w", err))
	}
	d.style = string(styleBytes)
}
