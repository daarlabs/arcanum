package form

import (
	"net/http"
)

type Builder struct {
	fields      []*FieldBuilder
	request     *http.Request
	method      string
	action      string
	name        string
	contentType string
	limit       int
	submitted   bool
	hx          bool
	security    security
	messages    Messages
	state       state
}

const (
	DefaultBodyLimit = 256
)

func New(fields ...*FieldBuilder) *Builder {
	return &Builder{
		fields:   fields,
		limit:    DefaultBodyLimit,
		messages: defaultMessages,
	}
}

func (b *Builder) Action(action string) *Builder {
	b.action = action
	return b
}

func (b *Builder) Add(name string) *FieldBuilder {
	field := Add(name)
	b.fields = append(b.fields, field)
	return field
}

func (b *Builder) Csrf(name, token string) *Builder {
	b.security = security{
		Enabled: len(name) > 0 && len(token) > 0,
		Name:    name,
		Token:   token,
	}
	return b
}

func (b *Builder) Messages(messages Messages) *Builder {
	if len(messages.Invalid) > 0 {
		b.messages.Invalid = messages.Invalid
	}
	if len(messages.MinText) > 0 {
		b.messages.MinText = messages.MinText
	}
	if len(messages.MaxText) > 0 {
		b.messages.MaxText = messages.MaxText
	}
	if len(messages.MinNumber) > 0 {
		b.messages.MinNumber = messages.MinNumber
	}
	if len(messages.MaxNumber) > 0 {
		b.messages.MaxNumber = messages.MaxNumber
	}
	if len(messages.Multipart) > 0 {
		b.messages.Multipart = messages.Multipart
	}
	if len(messages.Required) > 0 {
		b.messages.Required = messages.Required
	}
	if len(messages.Email) > 0 {
		b.messages.Email = messages.Email
	}
	return b
}

func (b *Builder) Get(name string) *FieldBuilder {
	for _, f := range b.fields {
		if f.name != name {
			continue
		}
		return f
	}
	return nil
}

func (b *Builder) Limit(limit int) *Builder {
	b.limit = limit
	return b
}
func (b *Builder) Method(method string) *Builder {
	b.method = method
	return b
}

func (b *Builder) Name(name string) *Builder {
	b.name = name
	return b
}

func (b *Builder) Request(request *http.Request) *Builder {
	b.request = request
	return b
}

func (b *Builder) Hx() *Builder {
	b.hx = true
	return b
}

func (b *Builder) State(state state) *Builder {
	b.state = state
	return b
}

func (b *Builder) isValid() bool {
	if !b.submitted {
		return true
	}
	for _, field := range b.fields {
		if !field.valid {
			return false
		}
	}
	return true
}
