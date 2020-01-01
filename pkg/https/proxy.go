package https

import (
	"net/http"

	httpproxy "github.com/huazhihao/scooter/pkg/http"
	"github.com/huazhihao/scooter/pkg/log"
)

// NewProxy creates a new tcp proxy
func NewHttpsProxy(p HttpsProxy) (*HttpsProxy, error) {
	hp, err := httpproxy.NewHttpProxy(p.HttpProxy)
	if err != nil {
		return nil, err
	}
	p.HttpProxy = *hp
	return &p, nil
}

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (p *HttpsProxy) ListenAndServe() {
	server := http.NewServeMux()
	server.HandleFunc("/", p.ServeHTTP)
	log.Debugf("Handling https connection on %s", p.Address)
	err := http.ListenAndServeTLS(p.Address, p.TLS.Cert, p.TLS.Key, server)
	if err != nil {
		log.Fatalf("Error while listening https connection on %s: %v", p.Address, err)
	}
}
