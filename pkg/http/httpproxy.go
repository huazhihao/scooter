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

package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/huazhihao/scooter/pkg/log"
)

// NewHTTPProxy creates a new tcp proxy
func NewHTTPProxy(p HTTPProxy) (*HTTPProxy, error) {
	err := p.reload()
	if err != nil {
		return nil, err
	}
	return &p, nil
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

func (p *HTTPProxy) getLongestMatchingRule(path string) int {
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

// reload reloads settings during runtime
func (p *HTTPProxy) reload() error {
	for i, r := range p.Rules {
		{
			url, err := url.Parse(r.URL)
			if err != nil {
				return err
			}
			log.Debugf("set rule mapping %s=>%s", r.Path, url)
			p.Rules[i].url = url
		}
	}
	return nil
}

func (p *HTTPProxy) delegateDirector(req *http.Request) {
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
func (p *HTTPProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Debugf("receiving %s request from %s/%s", req.Method, req.Host, req.URL.Path)

	r := &httputil.ReverseProxy{
		Director: p.delegateDirector,
	}

	r.ServeHTTP(w, req)
}

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (p *HTTPProxy) ListenAndServe() {
	server := http.NewServeMux()
	server.HandleFunc("/", p.ServeHTTP)
	log.Debugf("Handling http connection on %s", p.Address)
	err := http.ListenAndServe(p.Address, server)
	if err != nil {
		log.Fatalf("Error while listening http connection on %s: %v", p.Address, err)
	}
}
