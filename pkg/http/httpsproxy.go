package http

import (
	"net/http"

	"github.com/huazhihao/scooter/pkg/log"
)

// NewProxy creates a new tcp proxy
func NewHttpsProxy(p HttpsProxy) (*HttpsProxy, error) {
	_, err := NewHttpProxy(p.HttpProxy)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (p *HttpsProxy) ListenAndServe() {
	server := http.NewServeMux()
	server.HandleFunc("/", p.HttpProxy.ServeHTTP)
	log.Debugf("Handling https connection on %s", p.HttpProxy.Address)
	err := http.ListenAndServeTLS(p.HttpProxy.Address, p.TLS.Cert, p.TLS.Key, server)
	if err != nil {
		log.Fatalf("Error while listening https connection on %s: %v", p.HttpProxy.Address, err)
	}
}
