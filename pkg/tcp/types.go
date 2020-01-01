package tcp

import (
	"net"
	"time"
)

type TcpProxy struct {
	Name    string
	Address string
	Remote  string

	remote   *net.TCPAddr
	deadline time.Duration
}
