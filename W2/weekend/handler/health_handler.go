package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"w2/weekend/entity"
	"w2/weekend/service"
)

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

func collectHealthStatus() []entity.ServerStatus {
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

	return res
}
