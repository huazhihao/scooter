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

package metrics

import (
	"github.com/huazhihao/scooter/pkg/http"
)

// MetricsServer defines a MetricsServer data structure
type MetricsServer struct {
	Prometheus //TODO
}

// Prometheus defines the Prometheus interface of MetricsServer
type Prometheus struct {
	Address string
	TLS     http.TLS
}
