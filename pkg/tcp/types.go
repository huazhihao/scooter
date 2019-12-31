package tcp

import (
	"net"
	"time"
)

type TcpProxy struct {
	Name     string
	Bind     string
	Protocol string
	Remote   string

	remote   *net.TCPAddr
	deadline time.Duration
}
