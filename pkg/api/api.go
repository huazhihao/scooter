package api

import (
	"encoding/json"
	"net/http"

	httpproxy "github.com/huazhihao/scooter/pkg/http"
	"github.com/huazhihao/scooter/pkg/log"
)

type ApiServer struct {
	Address string
	TLS     httpproxy.TLS
}

func (s *ApiServer) stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	output := map[string]interface{}{
		"status": "ok",
	}
	json.NewEncoder(w).Encode(output)
}

func (s *ApiServer) dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Dashboard"))
}

// ListenAndServe listens on proxy.Address and then calls Serve to handle
// requests on incoming connections.
func (s *ApiServer) ListenAndServe() {
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
