package hx

import "github.com/daarlabs/arcanum/gox"

func Sync(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-sync")(value...)
}
