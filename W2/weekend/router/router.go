package router

import (
	"net/http"
	"w2/weekend/handler"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Mux *http.ServeMux = http.NewServeMux()

func init() {
	Mux.HandleFunc("GET /health", handler.HealthHandler)
	Mux.Handle("GET /metrics", promhttp.Handler())
	Mux.HandleFunc("GET /live", handler.LiveHandler)
	Mux.HandleFunc("GET /check-config", handler.ConfigCheckHandler)
}
