package metrics

import (
	"github.com/huazhihao/scooter/pkg/https"
)

type Prometheus struct {
	Address string
	TLS     https.TLS
}

type MetricsServer struct {
	Prometheus
}
