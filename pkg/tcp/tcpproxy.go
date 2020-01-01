// Copyright © 2019 Hua Zhihao <ihuazhihao@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcp

import (
	"io"
	"net"
	"time"

	"github.com/huazhihao/scooter/pkg/log"
)

// NewTCPProxy creates a new tcp proxy
func NewTCPProxy(p TCPProxy) (*TCPProxy, error) {
	remote, err := net.ResolveTCPAddr("tcp", p.Remote)
	if err != nil {
		return nil, err
	}
	p.remote = remote
	return &p, nil
}

// ServeTCP forwards the connection to a backend
func (p *TCPProxy) ServeTCP(conn net.Conn) {
	log.Debugf("Handling tcp connection from %s", conn.RemoteAddr())

	defer conn.Close()

	relay, err := net.DialTCP("tcp", nil, p.remote)
	if err != nil {
		log.Errorf("Error while connection to backend: %v", err)
		return
	}
	defer relay.Close()

	errChan := make(chan error)
	go p.connCopy(conn, relay, errChan)
	go p.connCopy(relay, conn, errChan)

	err = <-errChan
	if err != nil {
		log.Errorf("Error during connection: %v", err)
	}

	<-errChan
}

func (p TCPProxy) connCopy(dst, src net.Conn, errCh chan error) {
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

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (p *TCPProxy) ListenAndServe() {
	log.Debugf("Handling tcp connection on %s", p.Address)
	ln, err := net.Listen("tcp", p.Address)
	if err != nil {
		log.Fatalf("Error while listening tcp connection on %s: %v", p.Address, err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Errorf("Error while accepting connection from %v: %v", conn.RemoteAddr(), err)
			return
		}
		go p.ServeTCP(conn)
	}
}
