package https

import (
	"github.com/huazhihao/scooter/pkg/http"
)

type HttpsProxy struct {
	http.HttpProxy
	TLS
}

type TLS struct {
	Cert string
	Key  string
}
