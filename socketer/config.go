package socketer

type Config struct {
	ReadLimit  int64
	WriteLimit int64

	ReadBufferSize  int
	WriteBufferSize int
}

const (
	defaultReadLimit  = 512
	defaultWriteLimit = 512

	defaultReadBufferSize  = 1024
	defaultWriteBufferSize = 1024
)

var (
	defaultConfig = Config{
		ReadLimit:       defaultReadLimit,
		WriteLimit:      defaultWriteLimit,
		ReadBufferSize:  defaultReadBufferSize,
		WriteBufferSize: defaultWriteBufferSize,
	}
)
