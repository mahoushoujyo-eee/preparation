package weekend

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"w2/weekend/config"
	"w2/weekend/router"
)

func RunHealthCollector() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           router.Mux,
		ReadTimeout:       parseDuration(cfg.Server.ReadTimeout, 5*time.Second),
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      parseDuration(cfg.Server.WriteTimeout, 10*time.Second),
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Printf("Starting health collector on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}

// 空或非法时退回默认值，避免配置写错就崩。
func parseDuration(s string, def time.Duration) time.Duration {
	if s == "" {
		return def
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Printf("invalid duration %q, fallback to %s: %v", s, def, err)
		return def
	}
	return d
}
