package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/huazhihao/scooter/pkg/log"
)

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

func (p *Proxy) getLongestMatchingRule(path string) int {
	maxlen := -1
	idx := -1
	for i, r := range p.Rules {
		if strings.HasPrefix(path, r.Path) {
			l := len(r.Path)
			if l > maxlen {
				maxlen = l
				idx = i
			}
		}
	}
	return idx
}

// Reload reloads settings during runtime
func (p *Proxy) Reload() error {
	for i, r := range p.Rules {
		{
			url, err := url.Parse(r.Url)
			if err != nil {
				return err
			}
			log.Debugf("set rule mapping %s=>%#v", r.Path, url)
			p.Rules[i].url = url
		}
		p.Rules[i].urls = []*url.URL{}
		for _, u := range r.Urls {
			url, err := url.Parse(u)
			if err != nil {
				return err
			}
			p.Rules[i].urls = append(p.Rules[i].urls, url)
		}
	}
	return nil
}

func (p *Proxy) delegateDirector(req *http.Request) {
	i := p.getLongestMatchingRule(req.URL.Path)

	if i >= 0 {
		rule := p.Rules[i]
		target := rule.url
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
		// TODO: error handling
	}

	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", "")
	}
}

// ServeHTTP receives and handles a single http request
func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Debugf("receiving %s request from %s/%s", req.Method, req.Host, req.URL.Path)

	r := &httputil.ReverseProxy{
		Director: p.delegateDirector,
	}

	r.ServeHTTP(rw, req)
}

// ListenAndServe listens on proxy.bind and then calls Serve to handle
// requests on incoming connections.
func (p *Proxy) ListenAndServe() {
	http.HandleFunc("/", p.ServeHTTP)
	log.Debugf("Handling HTTP connection on %s", p.Bind)
	server := &http.Server{Addr: p.Bind}
	err := server.ListenAndServe()
	if err != nil {
		log.Debugf("Error while listening http connection: %v", err)
	}
}
