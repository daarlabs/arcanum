package hx

import "github.com/daarlabs/arcanum/gox"

func Trigger(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-trigger")(value...)
}
