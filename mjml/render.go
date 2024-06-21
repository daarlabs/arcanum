package mjml

import (
	"context"
	
	"github.com/Boostport/mjml-go"
	
	"github.com/daarlabs/arcanum/gox"
)

func Render(nodes ...gox.Node) (string, error) {
	return mjml.ToHTML(context.Background(), gox.Render(nodes...))
}

func MustRender(nodes ...gox.Node) string {
	html, err := Render(nodes...)
	if err != nil {
		panic(err)
	}
	return html
}
