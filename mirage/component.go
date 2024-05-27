package mirage

import (
	"encoding/json"
	"reflect"
	"strings"
	
	"github.com/daarlabs/arcanum/gox"
)

type MandatoryComponent interface {
	gox.Node
	Name() string
	Mount()
}

type Component struct {
	Ctx `json:"-"`
}

type component struct {
	ct     MandatoryComponent
	ctx    *ctx
	v      reflect.Value
	t      reflect.Type
	route  *Route
	action string
}

type componentCtx struct {
	name string
}

var (
	componentType = reflect.TypeOf(Component{})
)

func createComponent(ct MandatoryComponent, ctx *ctx, route *Route, action string) *component {
	c := &component{
		ct:     ct,
		ctx:    ctx,
		t:      reflect.TypeOf(ct),
		v:      reflect.ValueOf(ct),
		route:  route,
		action: action,
	}
	return c
}

func (c *component) render() gox.Node {
	var methodName string
	var shouldCallAction bool
	isActionRequest := len(c.action) > 0
	if isActionRequest {
		methodName, shouldCallAction = c.getAction()
	}
	if isActionRequest && !shouldCallAction {
		return gox.Fragment()
	}
	if !isActionRequest || (isActionRequest && shouldCallAction) {
		c.mustGet()
		c.injectContext()
		c.ct.Mount()
	}
	if isActionRequest && shouldCallAction {
		c.callAction(methodName)
		return gox.Fragment()
	}
	return c.ct.Node()
}

func (c *component) getAction() (string, bool) {
	action := c.action
	parts := strings.Split(action, namePrefixDivider)
	if len(parts) < 3 {
		return "", false
	}
	n := len(parts)
	compName := parts[n-2]
	methodName := parts[n-1]
	return methodName, compName == c.ct.Name() && c.v.MethodByName(methodName).IsValid()
}

func (c *component) callAction(methodName string) {
	method := c.v.MethodByName(methodName)
	if !method.IsValid() {
		return
	}
	methodResult := method.Call([]reflect.Value{})
	c.save()
	if len(methodResult) == 0 {
		return
	}
	*c.ctx.write = false
	switch r := methodResult[0].Interface().(type) {
	case error:
		c.ctx.err = r
	}
}

func (c *component) injectContext() {
	compField := c.v.Elem().FieldByName(componentType.Name())
	if !compField.IsValid() {
		return
	}
	compCtx := *c.ctx
	compCtx.component = &componentCtx{
		name: c.ct.Name(),
	}
	compField.Set(reflect.ValueOf(Component{Ctx: &compCtx}))
}

func (c *component) get() error {
	ct, ok := c.ctx.state.Components[c.route.Name+namePrefixDivider+c.ct.Name()]
	if !ok {
		return nil
	}
	bytes, err := json.Marshal(ct)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &c.ct)
}

func (c *component) mustGet() {
	err := c.get()
	if err != nil {
		panic(err)
	}
}

func (c *component) save() {
	c.ctx.state.Components[c.route.Name+namePrefixDivider+c.ct.Name()] = c.ct
	c.ctx.state.mustSave()
}
