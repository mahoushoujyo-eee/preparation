package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"w2/weekend/config"
	"w2/weekend/entity"
	"w2/weekend/service"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res, statusCode := collectHealthStatus()
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("marshal failed %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(bytes); err != nil {
		log.Printf("write failed %v", err)
	}
}

func LiveHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("living")); err != nil {
		log.Printf("write failed %v", err)
	}
}

func ConfigCheckHandler(w http.ResponseWriter, r *http.Request){
	config := config.Get()
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(config)
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

func collectHealthStatus() ([]entity.ServerStatus, int) {
	res := make([]entity.ServerStatus, 2, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		res[0] = service.CheckUserServiceHealth()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		res[1] = service.CheckQuotableServiceHealth()
		wg.Done()
	}()

	wg.Wait()

	cfg := config.Get()
	for _, val := range res{
    critical := cfg != nil && cfg.Services[val.ServiceName].Critical
		if critical && val.StatusCode != http.StatusOK {
			return res, http.StatusServiceUnavailable
		}
	}

	return res, http.StatusOK
}
