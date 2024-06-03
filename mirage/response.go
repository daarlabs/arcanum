package mirage

import (
	"net/url"
	
	"github.com/daarlabs/arcanum/gox"
	
	"github.com/daarlabs/arcanum/sender"
)

type Response interface {
	sender.ExtendableSend
	Status(statusCode int) Response
	Refresh() error
	Layout(name string) Response
	Render(nodes ...gox.Node) error
	Intercept() Intercept
}

type response struct {
	*sender.Sender
	ctx    *ctx
	layout *layout
	l      layoutFactory
}

func (r *response) Refresh() error {
	if !r.ctx.Request().Is().Action() {
		return r.Redirect(r.ctx.Generate().Current())
	}
	path, err := url.JoinPath(r.ctx.Config().Router.Prefix.Proxy, r.ctx.Request().Path())
	if err != nil {
		return r.Error(err)
	}
	return r.Redirect(path)
}

func (r *response) Layout(name string) Response {
	if r.layout == nil {
		return r
	}
	l, ok := r.layout.factories[name]
	if !ok {
		panic(ErrorInvalidLayout)
	}
	r.l = l
	return r
}

func (r *response) Status(statusCode int) Response {
	r.StatusCode = statusCode
	return r
}

func (r *response) Render(nodes ...gox.Node) error {
	if r.layout != nil && r.l != nil && !r.ctx.Request().Is().Hx() {
		return r.Html(gox.Render(r.l(r.ctx, nodes...)))
	}
	var style string
	rendered := gox.Render(nodes...)
	if r.ctx.tempest.Updated && r.ctx.Request().Is().Hx() {
		style = gox.Render(gox.Style(gox.Element(), gox.Type("text/css"), gox.Raw(r.ctx.tempest.Build())))
	}
	return r.Html(style + rendered)
}

func (r *response) Intercept() Intercept {
	return interceptor{
		Sender: r.Sender,
		err:    r.ctx.err,
	}
}
