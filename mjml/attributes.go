package mjml

import "github.com/daarlabs/arcanum/gox"

func Alt(value ...string) gox.Node {
	return gox.CreateAttribute[string]("alt")(value...)
}

func BaseUrl(value ...string) gox.Node {
	return gox.CreateAttribute[string]("base-url")(value...)
}

func Dir(value ...string) gox.Node {
	return gox.CreateAttribute[string]("dir")(value...)
}

func Href(value ...string) gox.Node {
	return gox.CreateAttribute[string]("href")(value...)
}

func Inline(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inline")(value...)
}

func Lang(value ...string) gox.Node {
	return gox.CreateAttribute[string]("lang")(value...)
}

func Name(value ...string) gox.Node {
	return gox.CreateAttribute[string]("name")(value...)
}

func Owa(value ...string) gox.Node {
	return gox.CreateAttribute[string]("owa")(value...)
}

func Rel(value ...string) gox.Node {
	return gox.CreateAttribute[string]("rel")(value...)
}

func Role(value ...string) gox.Node {
	return gox.CreateAttribute[string]("role")(value...)
}

func Src(value ...string) gox.Node {
	return gox.CreateAttribute[string]("src")(value...)
}

func SrcSet(value ...string) gox.Node {
	return gox.CreateAttribute[string]("srcset")(value...)
}

func Target(value ...string) gox.Node {
	return gox.CreateAttribute[string]("target")(value...)
}

func Thumbnails(value ...string) gox.Node {
	return gox.CreateAttribute[string]("thumbnails")(value...)
}

func ThumbnailsSrc(value ...string) gox.Node {
	return gox.CreateAttribute[string]("thumbnails-src")(value...)
}

func UseMap(value ...string) gox.Node {
	return gox.CreateAttribute[string]("usemap")(value...)
}
