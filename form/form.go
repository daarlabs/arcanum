package form

import (
	"github.com/daarlabs/arcanum/alpine"
	"github.com/daarlabs/arcanum/gox"
)

type Form struct {
	Security    security
	Method      string
	ContentType string
	Action      string
	Valid       bool
	Submitted   bool
	Hx          bool
}

func (f Form) Csrf() gox.Node {
	return Csrf(f.Security.Name, f.Security.Token)
}

func (f Form) Node(nodes ...gox.Node) gox.Node {
	return gox.Form(
		gox.Method(f.Method),
		gox.Action(f.Action),
		gox.EncType(f.ContentType),
		gox.If(f.Security.Enabled, Csrf(f.Security.Name, f.Security.Token)),
		gox.Fragment(nodes...),
		alpine.Data(map[string]any{"submitting": false}),
		alpine.Class(map[string]string{"form-request": "submitting"}),
		alpine.Submit("submitting = true"),
	)
}
