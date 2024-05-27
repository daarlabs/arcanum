package hx

import "github.com/daarlabs/arcanum/gox"

func Indicator(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-indicator")(value...)
}
