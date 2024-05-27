package form

import "github.com/daarlabs/arcanum/gox"

const (
	CsrfName  = "__csrf_name__"
	CsrfToken = "__csrf_token__"
)

func Csrf(name, token string) gox.Node {
	return gox.Fragment(
		gox.Input(gox.Type("hidden"), gox.Name(CsrfName), gox.Value(name)),
		gox.Input(gox.Type("hidden"), gox.Name(CsrfToken), gox.Value(token)),
	)
}
