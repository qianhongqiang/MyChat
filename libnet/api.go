package libnet

import (
	"io"
	"net"
)

type Protocol interface {
	NewCodec(rw io.ReadWriter) Codec
}

type Codec interface {
	Receive() ([]byte, error)
	Send(interface{}) error
	Close() error
}

func Serve(network, address string, protocol Protocol, sendChanSize int) (*Server, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return NewServer(listener, protocol, sendChanSize), nil
}