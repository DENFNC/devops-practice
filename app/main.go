package main

import (
	"log"
	"net/http"

	"k8s-app/app/config"
	"k8s-app/app/handlers"
	"k8s-app/app/middleware"
	"k8s-app/app/services"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Создаем сервисы
	appService := services.NewAppService()

	// Создаем handlers
	handler := handlers.NewHandler(appService)

	// Запускаем инициализацию в отдельной горутине
	go appService.Initialize()

	// Настраиваем маршруты
	mux := http.NewServeMux()

	// Применяем middleware
	loggedMux := middleware.LoggingMiddleware(mux)

	// Регистрируем маршруты
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/ready", handler.ReadyHandler)
	mux.HandleFunc("/metrics", handler.MetricsHandler)
	mux.HandleFunc("/load", handler.LoadHandler)

	log.Printf("Запускаем приложение на порту %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, loggedMux))
}
