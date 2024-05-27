package mjml

import "github.com/daarlabs/arcanum/gox"

func Attributes(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-attributes")(nodes...)
}

func Breakpoint(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-breakpoint")(nodes...)
}

func Font(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-font")(nodes...)
}

func HtmlAttributes(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-html-attributes")(nodes...)
}

func Preview(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-preview")(nodes...)
}

func Style(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-style")(nodes...)
}

func Class(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-class")(nodes...)
}

func All(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-all")(nodes...)
}
