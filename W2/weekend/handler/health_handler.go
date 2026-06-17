package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ServerStatus struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	StatusCode  int    `json:"status_code"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res := collectHealthStatus()
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("marshal failed %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytes); err != nil {
		log.Printf("write failed %v", err)
	}
}

func collectHealthStatus() []ServerStatus {
	res := make([]ServerStatus, 0, 10)
	resChan := make(chan ServerStatus, 10)
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			resChan <- ServerStatus{
				ServiceName: fmt.Sprintf("service %d", i),
				Status:      "healthy",
				StatusCode:  1,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for val := range resChan {
		res = append(res, val)
	}

	return res
}
