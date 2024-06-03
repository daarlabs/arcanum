package tempest

import "strings"

type Context struct {
	Tempest  *Tempest
	builders []*Builder
	classes  map[string]string
	Updated  bool
}

func (c *Context) Class() *Builder {
	b := &Builder{
		Context: c,
		classes: make([]string, 0),
	}
	c.builders = append(c.builders, b)
	return b
}

func (c *Context) Add(k, v string) {
	c.classes[k] = v
	c.Tempest.classes[k] = v
	c.Updated = true
}

func (c *Context) Build() string {
	result := make([]string, len(c.classes))
	i := 0
	for _, v := range c.classes {
		result[i] = v
		i++
	}
	return strings.Join(result, " ")
}
