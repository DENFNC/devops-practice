package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"k8s-app/app/models"
	"k8s-app/app/services"
)

// Handler содержит зависимости для обработчиков
type Handler struct {
	appService *services.AppService
}

// NewHandler создает новый экземпляр Handler
func NewHandler(appService *services.AppService) *Handler {
	return &Handler{
		appService: appService,
	}
}

// HomeHandler обрабатывает главную страницу
func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := h.appService.GetStatus()

	if !status.Initialized {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(models.Response{
			Status:    "initializing",
			Message:   "Приложение инициализируется, подождите...",
			Data:      map[string]interface{}{"uptime": status.Uptime.Seconds()},
			Timestamp: time.Now(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Status:  "ready",
		Message: "Приложение готово к работе",
		Data: map[string]interface{}{
			"uptime":          status.Uptime.Seconds(),
			"memory_usage_mb": h.appService.GetMemoryUsage(),
		},
		Timestamp: time.Now(),
	})
}

// HealthHandler обрабатывает health check
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := h.appService.GetStatus()

	if !status.Initialized {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(models.Response{
			Status:    "not_ready",
			Timestamp: time.Now(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.HealthResponse{
		Status:                 "healthy",
		InitializationComplete: true,
		Uptime:                 status.Uptime.Seconds(),
		Timestamp:              time.Now(),
	})
}

// ReadyHandler обрабатывает readiness probe
func (h *Handler) ReadyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if !h.appService.IsInitialized() {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]bool{"ready": false})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"ready": true})
}

// MetricsHandler возвращает метрики приложения
func (h *Handler) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := h.appService.GetStatus()

	json.NewEncoder(w).Encode(models.MetricsResponse{
		CPUPercent:             h.appService.GetCPUUsage(),
		MemoryMB:               h.appService.GetMemoryUsage(),
		UptimeSeconds:          status.Uptime.Seconds(),
		InitializationComplete: status.Initialized,
		Timestamp:              time.Now(),
	})
}

// LoadHandler имитирует обработку нагрузки
func (h *Handler) LoadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := h.appService.ProcessLoad()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(models.Response{
			Status:    "error",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Status:  "success",
		Message: "Request processed",
		Data: map[string]interface{}{
			"timestamp":       time.Now().Unix(),
			"memory_usage_mb": h.appService.GetMemoryUsage(),
		},
		Timestamp: time.Now(),
	})
}
