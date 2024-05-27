package filesystem

import (
	"context"
)

type Client interface {
	GetList() ([]string, error)
	Read(path string) ([]byte, error)
	Create(path string, data []byte) error
	Remove(path string) error
	
	MustGetList() []string
	MustRead(path string) []byte
	MustCreate(path string, data []byte)
	MustRemove(path string)
}

const (
	Local = "local"
	Cloud = "cloud"
)

func New(ctx context.Context, config Config) Client {
	switch config.Driver {
	case Local:
		return createLocal(config.Dir)
	case Cloud:
		if config.Cloud == nil {
			panic(ErrorMissingCloud)
		}
		return createCloud(ctx, config.Dir, config.Name, config.Cloud)
	}
	return nil
}
