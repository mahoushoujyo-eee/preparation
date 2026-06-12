package weekend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type ServerStatus struct{
	ServeName string `json:"service"`
	Status string `json:"status"`
	StatusCode int `json:"status_code"`
}

func RunHealthCollector() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		res := CollectHealthStatus()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

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
			semaphore <- struct{}{}
			resChan <- ServerStatus{
				ServeName: fmt.Sprintf("server%d", i),
				Status: "healthy",
				StatusCode: 1,
			}
			defer func ()  {
				wg.Done()
				<-semaphore
			}()
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