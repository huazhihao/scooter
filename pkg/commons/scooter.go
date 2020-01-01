package commons

import (
	"log"

	"github.com/huazhihao/scooter/pkg/http"
	"github.com/huazhihao/scooter/pkg/https"
	"github.com/huazhihao/scooter/pkg/tcp"
)

type Scooter struct {
	config Config
}

func NewScooter(config Config) *Scooter {
	f := Scooter{config: config}
	return &f
}

func (s *Scooter) Run() {
	done := make(chan bool)

	for _, p := range s.config.HttpProxies {
		proxy, err := http.NewHttpProxy(p)
		if err != nil {
			log.Fatalf("Unable to create http proxy: %v", err)
		}
		go proxy.ListenAndServe()

	}

	for _, p := range s.config.HttpsProxies {
		proxy, err := https.NewHttpsProxy(p)
		if err != nil {
			log.Fatalf("Unable to create https proxy: %v", err)
		}
		go proxy.ListenAndServe()
	}

	for _, p := range s.config.TcpProxies {
		proxy, err := tcp.NewTcpProxy(p)
		if err != nil {
			log.Fatalf("Unable to create tcp proxy: %v", err)
		}
		go proxy.ListenAndServe()
	}

	apiServer := &s.config.ApiServer
	go apiServer.ListenAndServe()

	<-done
}
