package entity

type ServerStatus struct {
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	StatusCode  int    `json:"status_code"`
	LatencyMs   int64  `json:"latency_ms"`
}

type MockUser struct {
	ID   int
	Name string
}

type Quotable struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}
