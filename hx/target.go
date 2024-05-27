package hx

import "github.com/daarlabs/arcanum/gox"

func Target(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-target")(value...)
}
