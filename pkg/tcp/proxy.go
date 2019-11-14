package tcp

import (
	"io"
	"net"
	"time"

	"github.com/huazhihao/scooter/pkg/log"
)

// Proxy forwards a TCP request to a TCP backend
type Proxy struct {
	addr     *net.TCPAddr
	deadline time.Duration
}

// NewProxy creates a new tcp proxy
func NewProxy(addr string, deadline time.Duration) (*Proxy, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Proxy{addr: tcpAddr, deadline: deadline}, nil
}

// ServeTCP forwards the connection to a backend
func (p *Proxy) ServeTCP(conn net.Conn) {
	log.Debugf("Handling tcp connection from %s", conn.RemoteAddr())

	defer conn.Close()

	connBackend, err := net.DialTCP("tcp", nil, p.addr)
	if err != nil {
		log.Errorf("Error while connection to backend: %v", err)
		return
	}
	defer connBackend.Close()

	errChan := make(chan error)
	go p.connCopy(conn, connBackend, errChan)
	go p.connCopy(connBackend, conn, errChan)

	err = <-errChan
	if err != nil {
		log.Errorf("Error during connection: %v", err)
	}

	<-errChan
}

func (p Proxy) connCopy(dst, src net.Conn, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err

	errClose := dst.Close()
	if errClose != nil {
		log.Debugf("Error while terminating connection: %v", errClose)
		return
	}

	if p.deadline >= 0 {
		err := dst.SetReadDeadline(time.Now().Add(p.deadline))
		if err != nil {
			log.Debugf("Error while setting deadline: %v", err)
		}
	}
}
