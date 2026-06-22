package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// MockServiceCallTotal 记录每个 mock service 被调用的总次数。
// 标签 service 用来区分是哪个 service 被调用了。
//
// 在 /metrics 输出形如：
//
//	mock_service_call_total{service="user_service"} 5
var MockServiceCallTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "mock_service_call_total",
		Help: "Total number of times each mock service was called.",
	},
	[]string{"service"},
)
