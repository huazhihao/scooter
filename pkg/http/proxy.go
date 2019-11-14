package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/huazhihao/scooter/pkg/log"
)

// Proxy forwards a HTTP request to a HTTP backend
type Proxy struct {
	mappings map[string]*url.URL
	Addr     string
}

// NewProxy creates a new http proxy
func NewProxy(addr string, mappings map[string]string) (*Proxy, error) {
	proxy := &Proxy{Addr: addr}
	err := proxy.UpdateMappings(mappings)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func joinURLPath(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func (s *Proxy) delegateDirector(req *http.Request) {
	target, ok := s.mappings[req.Host]
	if ok {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = joinURLPath(target.Path, req.URL.Path)
		req.Host = target.Host
		if target.RawQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
		}
	} else {

	}

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", "")
	}
}

// ServeHTTP receives and handles a single http request
func (s *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Debugf("receiving %s request from %s/%s", req.Method, req.Host, req.URL.Path)

	p := &httputil.ReverseProxy{
		Director: s.delegateDirector,
	}

	p.ServeHTTP(rw, req)
}

// UpdateMappings updated the mapping settings during runtime
func (s *Proxy) UpdateMappings(mappings map[string]string) error {
	s.mappings = map[string]*url.URL{}
	for k, v := range mappings {
		target, err := url.Parse(v)
		if err != nil {
			return err
		}
		s.mappings[k] = target
		log.Debugf("updating mapping: %s=>%v", k, target)
	}
	return nil
}

// ListenAndServe listens on proxy.Addr and then calls Serve to handle
// requests on incoming connections.
func (s *Proxy) ListenAndServe() {
	http.HandleFunc("/", s.ServeHTTP)
	log.Debugf("Handling HTTP connection on %s", s.Addr)
	server := &http.Server{Addr: s.Addr}
	err := server.ListenAndServe()
	if err != nil {
		log.Debugf("Error while listening http connection: %v", err)
	}
}
