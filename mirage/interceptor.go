package mirage

import "github.com/daarlabs/arcanum/sender"

type Intercept interface {
	Error() error
	Status() int
}

type interceptor struct {
	*sender.Sender
	err error
}

func (c interceptor) Error() error {
	return c.err
}

func (c interceptor) Status() int {
	return c.StatusCode
}
