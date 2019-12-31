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

// ListenAndServe listens on proxy.bind and then calls Serve to handle
// requests on incoming connections.
func (p *HttpsProxy) ListenAndServe() {
	http.HandleFunc("/", p.ServeHTTP)
	log.Debugf("Handling https connection on %s", p.Bind)
	err := http.ListenAndServeTLS(p.Bind, p.TLS.Cert, p.TLS.Key, nil)
	if err != nil {
		log.Debugf("Error while listening https connection: %v", err)
	}
}
