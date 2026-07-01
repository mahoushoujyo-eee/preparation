package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"w2/weekend/config"
	"w2/weekend/entity"
	"w2/weekend/metrics"
)

var (
	defaultQuotableUrl  string = "https://dummyjson.com/quotes/random"
	quotableServiceName string = "quotable_service"
)

// quotableURL 优先从配置读取；配置缺失或为空时退回默认值。
func quotableURL() string {
	if cfg := config.Get(); cfg != nil {
		if s, ok := cfg.Services[quotableServiceName]; ok && s.URL != "" {
			return s.URL
		}
	}
	return defaultQuotableUrl
}

func CheckQuotableServiceHealth() entity.ServerStatus {
	start := time.Now()
	_, err := GetOneQuotable()
	latency := time.Since(start).Milliseconds()
	if err != nil {
		log.Printf("[%s] health check failed: %v", quotableServiceName, err)
		return entity.ServerStatus{
			ServiceName: quotableServiceName,
			Status:      "unhealthy",
			StatusCode:  http.StatusBadGateway,
			LatencyMs:   latency,
		}
	}

	return entity.ServerStatus{
		ServiceName: quotableServiceName,
		Status:      "healthy",
		StatusCode:  http.StatusOK,
		LatencyMs:   latency,
	}
}

func GetOneQuotable() (*entity.Quotable, error) {
	metrics.MockServiceCallTotal.WithLabelValues(quotableServiceName).Inc()
	body, err := RequestWithTimeout(quotableURL(), "GET")
	if err != nil {
		return nil, err
	}

	var quotable entity.Quotable
	err = json.Unmarshal(body, &quotable)
	if err != nil {
		return nil, err
	}

	return &quotable, nil
}

func RequestWithTimeout(url string, method string) ([]byte, error) {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	resq, err := http.NewRequestWithContext(ctxWithTimeout, method, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	res, err := client.Do(resq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status: %s", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
