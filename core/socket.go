package core

type ISocket interface {
	Start() error
	Shutdown() error
	Read() (int, []byte, error)
	Write([]byte) error
	Close() error
}
