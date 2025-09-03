package models

import "time"

// Response представляет стандартный ответ API
type Response struct {
	Status    string      `json:"status"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// HealthResponse представляет ответ health check
type HealthResponse struct {
	Status                 string    `json:"status"`
	InitializationComplete bool      `json:"initialization_complete"`
	Uptime                 float64   `json:"uptime"`
	Timestamp              time.Time `json:"timestamp"`
}

// MetricsResponse представляет ответ с метриками
type MetricsResponse struct {
	CPUPercent             float64   `json:"cpu_percent"`
	MemoryMB               float64   `json:"memory_mb"`
	UptimeSeconds          float64   `json:"uptime_seconds"`
	InitializationComplete bool      `json:"initialization_complete"`
	Timestamp              time.Time `json:"timestamp"`
}

// LoadResponse представляет ответ на запрос нагрузки
type LoadResponse struct {
	Message       string  `json:"message"`
	Timestamp     int64   `json:"timestamp"`
	MemoryUsageMB float64 `json:"memory_usage_mb"`
}

// AppStatus представляет статус приложения
type AppStatus struct {
	Initialized bool
	StartTime   time.Time
	Uptime      time.Duration
}
