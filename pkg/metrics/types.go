package metrics

import (
	"github.com/huazhihao/scooter/pkg/http"
)

type Prometheus struct {
	Address string
	TLS     http.TLS
}

type MetricsServer struct {
	Prometheus
}
