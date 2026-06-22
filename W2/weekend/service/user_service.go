package service

import (
	"log"
	"time"
	"w2/weekend/dao"
	"w2/weekend/entity"
	"w2/weekend/metrics"
)

var userServiceName string = "user_service"

func CheckUserServiceHealth() entity.ServerStatus {
	metrics.MockServiceCallTotal.WithLabelValues(userServiceName).Inc()
	start := time.Now()
	err := dao.CheckDBCon()
	latency := time.Since(start).Milliseconds()
	if err != nil {
		log.Printf("[%s] health check failed: %v", userServiceName, err)
		return entity.ServerStatus{
			ServiceName: userServiceName,
			Status:      "unhealthy",
			StatusCode:  500,
			LatencyMs:   latency,
		}
	}

	return entity.ServerStatus{
		ServiceName: userServiceName,
		Status:      "healthy",
		StatusCode:  200,
		LatencyMs:   latency,
	}
}

func GetAllUsers() []entity.MockUser {
	metrics.MockServiceCallTotal.WithLabelValues(userServiceName).Inc()
	users, err := dao.QueryAllUsers()
	if err != nil {
		log.Printf("query users failed, %v", err)
	}

	return users
}
