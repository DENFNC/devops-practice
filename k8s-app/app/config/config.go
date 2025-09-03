package config

import (
"os"
"strconv"
"time"
)

// Config содержит конфигурацию приложения
type Config struct {
Port                string
LogLevel            string
MetricsEnabled      bool
InitializationDelay time.Duration
HealthCheckInterval time.Duration
ReadinessTimeout    time.Duration
LivenessTimeout     time.Duration
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
config := &Config{
Port:                getEnv("PORT", "8080"),
LogLevel:            getEnv("LOG_LEVEL", "info"),
MetricsEnabled:      getEnvBool("METRICS_ENABLED", true),
InitializationDelay: getEnvDuration("INIT_DELAY", 5*time.Second),
HealthCheckInterval: getEnvDuration("HEALTH_INTERVAL", 10*time.Second),
ReadinessTimeout:    getEnvDuration("READINESS_TIMEOUT", 3*time.Second),
LivenessTimeout:     getEnvDuration("LIVENESS_TIMEOUT", 5*time.Second),
}

return config
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
if value := os.Getenv(key); value != "" {
return value
}
return defaultValue
}

// getEnvBool получает булеву переменную окружения
func getEnvBool(key string, defaultValue bool) bool {
if value := os.Getenv(key); value != "" {
if parsed, err := strconv.ParseBool(value); err == nil {
return parsed
}
}
return defaultValue
}

// getEnvDuration получает переменную окружения как duration
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
if value := os.Getenv(key); value != "" {
if parsed, err := time.ParseDuration(value); err == nil {
return parsed
}
}
return defaultValue
}
