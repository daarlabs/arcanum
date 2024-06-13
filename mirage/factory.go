package mirage

import (
	"fmt"
	"net/http"
	
	"github.com/dchest/uniuri"
	
	"github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/hx"
	
	"github.com/daarlabs/arcanum/csrf"
	"github.com/daarlabs/arcanum/form"
)

type Factory interface {
	Component(ct MandatoryComponent) gox.Node
	Defer(link string, nodes ...gox.Node) gox.Node
	Form(fields ...*form.FieldBuilder) *form.Builder
}

type factory struct {
	ctx *ctx
}

func (f factory) Component(ct MandatoryComponent) gox.Node {
	var action string
	_ = f.ctx.Parse().Query(Action, &action)
	return createComponent(ct, f.ctx, f.ctx.route, action).render()
}

func (f factory) Defer(link string, nodes ...gox.Node) gox.Node {
	return gox.Div(
		hx.Get(link),
		hx.Trigger("load"),
		hx.Swap(hx.SwapOuterHtml),
		hx.Headers(Map{hx.RequestHeaderTrigger: "load"}),
		gox.Fragment(nodes...),
	)
}

func (f factory) Form(fields ...*form.FieldBuilder) *form.Builder {
	isCsrfEnabled := f.ctx.Config().Security.Csrf != nil && f.ctx.Config().Security.Csrf.IsEnabled()
	method := f.ctx.Request().Method()
	if f.ctx.Request().Is().Get() {
		method = http.MethodPost
	}
	link := f.ctx.Generate().Link(f.ctx.route.Name)
	if len(f.ctx.Request().Raw().URL.Query()) > 0 {
		q := make(Map)
		for k := range f.ctx.Request().Raw().URL.Query() {
			q[k] = f.ctx.Request().Raw().URL.Query().Get(k)
		}
		link += f.ctx.Generate().Query(q)
	}
	r := form.New(fields...).
		Limit(f.ctx.Config().Form.Limit).
		Method(method).
		Action(link).
		Request(f.ctx.r).
		Messages(
			form.Messages{
				Email:     f.ctx.Translate(f.ctx.config.Localization.Form.Email),
				Required:  f.ctx.Translate(f.ctx.config.Localization.Form.Required),
				MinText:   f.ctx.Translate(f.ctx.config.Localization.Form.MinText),
				MaxText:   f.ctx.Translate(f.ctx.config.Localization.Form.MaxText),
				MinNumber: f.ctx.Translate(f.ctx.config.Localization.Form.MinNumber),
				MaxNumber: f.ctx.Translate(f.ctx.config.Localization.Form.MaxNumber),
				Multipart: f.ctx.Translate(f.ctx.config.Localization.Form.Multipart),
				Invalid:   f.ctx.Translate(f.ctx.config.Localization.Form.Invalid),
			},
		)
	if isCsrfEnabled && !f.ctx.Request().Is().Action() {
		name := fmt.Sprintf("%s-%s", f.ctx.route.Name, uniuri.New())
		token := f.ctx.Csrf().MustCreate(
			csrf.Token{
				Name: name, UserAgent: f.ctx.Request().UserAgent(), Ip: f.ctx.Request().Ip(),
			},
		)
		r.Csrf(name, token)
	}
	return r
}
