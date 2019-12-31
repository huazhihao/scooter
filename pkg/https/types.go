package https

import (
	"github.com/huazhihao/scooter/pkg/http"
)

type HttpsProxy struct {
	http.HttpProxy
	TLS
}

type TLS struct {
	CA   string
	Cert string
	Key  string
}
