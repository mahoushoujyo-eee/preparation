package weekend

import (
	"log"
	"net/http"
	"time"

	"w2/weekend/router"
)

func RunHealthCollector() {
	srv := http.Server{
		Addr:              "127.0.0.1:8085",
		Handler:           router.Mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Printf("Starting health collector on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
