package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware логирует HTTP запросы
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Создаем wrapper для ResponseWriter для захвата статуса
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Выполняем следующий handler
		next.ServeHTTP(wrapped, r)

		// Логируем запрос
		duration := time.Since(start)
		log.Printf("%s %s %d %v %s",
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
			duration,
			r.UserAgent())
	})
}

// responseWriter обертка для http.ResponseWriter
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader перехватывает статус код
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
