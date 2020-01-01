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

package commons

import (
	"log"

	"github.com/huazhihao/scooter/pkg/http"
	"github.com/huazhihao/scooter/pkg/tcp"
)

// Scooter defines scooter runtime
type Scooter struct {
	config Config
}

// NewScooter creates a scooter runtime
func NewScooter(config Config) *Scooter {
	f := Scooter{config: config}
	return &f
}

// Run starts running scooter runtime
func (s *Scooter) Run() {
	done := make(chan bool)

	for _, p := range s.config.HTTPProxies {
		proxy, err := http.NewHTTPProxy(p)
		if err != nil {
			log.Fatalf("Unable to create http proxy: %v", err)
		}
		go proxy.ListenAndServe()

	}

	for _, p := range s.config.HTTPSProxies {
		proxy, err := http.NewHTTPSProxy(p)
		if err != nil {
			log.Fatalf("Unable to create https proxy: %v", err)
		}
		go proxy.ListenAndServe()
	}

	for _, p := range s.config.TCPProxies {
		proxy, err := tcp.NewTCPProxy(p)
		if err != nil {
			log.Fatalf("Unable to create tcp proxy: %v", err)
		}
		go proxy.ListenAndServe()
	}

	apiServer := &s.config.APIServer
	go apiServer.ListenAndServe()

	<-done
}
