package api

import (
	"github.com/huazhihao/scooter/pkg/https"
)

type Entrypoint struct {
	Address string
	TLS     https.TLS
}
