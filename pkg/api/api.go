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

package api

import (
	"encoding/json"
	"net/http"

	httpproxy "github.com/huazhihao/scooter/pkg/http"
	"github.com/huazhihao/scooter/pkg/log"
)

// APIServer stands for a api server data structure
type APIServer struct {
	Address string
	TLS     httpproxy.TLS
}

func (s *APIServer) stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output := map[string]interface{}{
		"status": "ok",
	}
	json.NewEncoder(w).Encode(output)
}

func (s *APIServer) dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Dashboard"))
}

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (s *APIServer) ListenAndServe() {
	server := http.NewServeMux()
	server.Handle("/", http.RedirectHandler("/dashboard", 307))
	server.HandleFunc("/dashboard", s.dashboard)
	server.HandleFunc("/stats", s.stats)

	if s.TLS.Cert != "" && s.TLS.Key != "" {
		log.Debugf("Handling https connection on %s", s.Address)
		err := http.ListenAndServeTLS(s.Address, s.TLS.Cert, s.TLS.Key, server)
		if err != nil {
			log.Debugf("Error while listening https connection: %v", err)
		}
	} else {
		log.Debugf("Handling http connection on %s", s.Address)
		err := http.ListenAndServe(s.Address, server)
		if err != nil {
			log.Fatalf("Error while listening http connection: %v", err)
		}
	}
}
