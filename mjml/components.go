package mjml

import "github.com/daarlabs/arcanum/gox"

func Mjml(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mjml")(nodes...)
}

func Head(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-head")(nodes...)
}

func Body(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-body")(nodes...)
}

func Include(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-include")(nodes...)
}
