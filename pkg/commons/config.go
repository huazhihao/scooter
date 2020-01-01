package commons

import (
	"github.com/huazhihao/scooter/pkg/api"
	"github.com/huazhihao/scooter/pkg/http"
	"github.com/huazhihao/scooter/pkg/https"
	"github.com/huazhihao/scooter/pkg/metrics"
	"github.com/huazhihao/scooter/pkg/tcp"
)

type Config struct {
	HttpProxies   []http.HttpProxy   `yaml:"http"`
	HttpsProxies  []https.HttpsProxy `yaml:"https"`
	TcpProxies    []tcp.TcpProxy     `yaml:"tcp"`
	ApiServer     api.Entrypoint     `yaml:"api"`
	MetricsServer metrics.Entrypoint `yaml:"metrics"`
}
