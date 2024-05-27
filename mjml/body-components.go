package mjml

import "github.com/daarlabs/arcanum/gox"

func Accordion(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-accordion")(nodes...)
}

func AccordionElement(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-accordion-element")(nodes...)
}

func AccordionTitle(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-accordion-title")(nodes...)
}

func AccordionText(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-accordion-text")(nodes...)
}

func Button(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-button")(nodes...)
}

func Carousel(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-carousel")(nodes...)
}

func CarouselImage(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-carousel-image")(nodes...)
}

func Column(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-column")(nodes...)
}

func Divider(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-divider")(nodes...)
}

func Group(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-group")(nodes...)
}

func Hero(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-hero")(nodes...)
}

func Image(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-image")(nodes...)
}

func NavBar(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-navbar")(nodes...)
}

func NavBarLink(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-navbar-link")(nodes...)
}

func Raw(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-raw")(nodes...)
}

func Section(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-section")(nodes...)
}

func Social(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-social")(nodes...)
}

func SocialElement(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-social-element")(nodes...)
}

func Spacer(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-spacer")(nodes...)
}

func Table(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-table")(nodes...)
}

func Text(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-text")(nodes...)
}

func Wrapper(nodes ...gox.Node) gox.Node {
	return gox.CreateElement("mj-wrapper")(nodes...)
}
