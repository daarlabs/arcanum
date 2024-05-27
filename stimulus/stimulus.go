package stimulus

import (
	"fmt"
	
	"github.com/daarlabs/arcanum/gox"
)

func Controller(name string) gox.Node {
	return gox.CreateAttribute[string]("data-controller")(name)
}

func Target(name string, value any) gox.Node {
	return gox.CreateAttribute[string](fmt.Sprintf(`data-%s-target`, name))(fmt.Sprintf("%v", value))
}

func Action(event, controller, method string) gox.Node {
	return gox.CreateAttribute[string]("data-action")(fmt.Sprintf("%s->%s#%s", event, controller, method))
}

func Value(controller, name string, value any) gox.Node {
	return gox.CreateAttribute[string](fmt.Sprintf("data-%s-%s-value", controller, name))(fmt.Sprintf("%v", value))
}

func Outlet(controller, outletController, selector string) gox.Node {
	return gox.CreateAttribute[string](fmt.Sprintf("data-%s-%s-outlet", controller, outletController))(selector)
}
