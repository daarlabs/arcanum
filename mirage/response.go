package mirage

import (
	"net/url"
	
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
	
	"github.com/daarlabs/arcanum/sender"
)

type Response interface {
	sender.ExtendableSend
	Status(statusCode int) Response
	Refresh() error
	Render(nodes ...gox.Node) error
	Intercept() Intercept
}

type response struct {
	*sender.Sender
	ctx   *ctx
	route *Route
}

func (r *response) Refresh() error {
	if !r.ctx.Request().Is().Action() {
		return r.Redirect(r.ctx.Generate().Current())
	}
	path, err := url.JoinPath("/", r.ctx.Config().Router.Prefix.Proxy, r.ctx.Request().Path())
	if err != nil {
		return r.Error(err)
	}
	return r.Redirect(path)
}

func (r *response) Status(statusCode int) Response {
	r.StatusCode = statusCode
	return r
}

func (r *response) Render(nodes ...gox.Node) error {
	isHx := r.ctx.Request().Is().Hx()
	if !r.ctx.Request().Is().Get() && !isHx {
		return r.ctx.Response().Redirect(r.ctx.Generate().Current())
	}
	if r.route.Layout != nil && !isHx {
		return r.Html(gox.Render(r.route.Layout(r.ctx, nodes...)))
	}
	if isHx {
		styles := tempest.NamedStyles(r.ctx.Request().Action())
		if len(styles) > 0 {
			return r.Html(
				gox.Render(
					gox.Fragment(nodes...),
					gox.Style(
						gox.Element(),
						gox.Type("text/css"),
						gox.Raw(styles),
					),
				),
			)
		}
	}
	return r.Html(gox.Render(nodes...))
}

func (r *response) Intercept() Intercept {
	return interceptor{
		Sender: r.Sender,
		err:    r.ctx.err,
	}
}
