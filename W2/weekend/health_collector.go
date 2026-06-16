package weekend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ServerStatus struct{
	ServiceName string `json:"service_name"`
	Status string `json:"status"`
	StatusCode int `json:"status_code"`
}

func RunHealthCollector() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		res := CollectHealthStatus()
		w.Header().Set("Content-Type", "application/json")
		bytes, err := json.Marshal(res)
		if err != nil {
			log.Printf("marshal failed %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(bytes)
		if err != nil {
			log.Printf("write failed %v", err)
		}
	})

	mux.Handle("GET /metrics", promhttp.Handler())

	srv := http.Server{
		Addr: "127.0.0.1:8085",
		Handler: mux,
		ReadTimeout: 5*time.Second,
        ReadHeaderTimeout: 3 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout: 30 * time.Second,
        MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting health collector on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}

func CollectHealthStatus() []ServerStatus{
	res := make([]ServerStatus, 0)
	resChan := make(chan ServerStatus, 10)
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() {<-semaphore}()
			resChan <- ServerStatus{
				ServiceName: fmt.Sprintf("service %d", i),
				Status: "healthy",
				StatusCode: 1,
			}
		}(i)
	}

	go func(){
		wg.Wait()
		close(resChan)
	}()

	for val := range resChan {
		res = append(res, val)
	}

	return res
}