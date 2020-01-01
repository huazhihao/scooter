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

	"github.com/huazhihao/scooter/pkg/log"
)

// NewHTTPSProxy creates a new tcp proxy
func NewHTTPSProxy(p HTTPSProxy) (*HTTPSProxy, error) {
	_, err := NewHTTPProxy(p.HTTPProxy)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (p *HTTPSProxy) ListenAndServe() {
	server := http.NewServeMux()
	server.HandleFunc("/", p.HTTPProxy.ServeHTTP)
	log.Debugf("Handling https connection on %s", p.HTTPProxy.Address)
	err := http.ListenAndServeTLS(p.HTTPProxy.Address, p.TLS.Cert, p.TLS.Key, server)
	if err != nil {
		log.Fatalf("Error while listening https connection on %s: %v", p.HTTPProxy.Address, err)
	}
}
