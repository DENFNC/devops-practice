package services

import (
	"log"
	"runtime"
	"sync"
	"time"

	"k8s-app/app/models"
)

// AppService управляет состоянием приложения
type AppService struct {
	status *models.AppStatus
	mu     sync.RWMutex
}

// NewAppService создает новый экземпляр AppService
func NewAppService() *AppService {
	return &AppService{
		status: &models.AppStatus{
			Initialized: false,
			StartTime:   time.Now(),
		},
	}
}

// Initialize запускает процесс инициализации приложения
func (s *AppService) Initialize() {
	log.Println("Начинаем инициализацию приложения...")

	// Имитируем интенсивную работу при инициализации
	for i := 0; i < 50; i++ { // ~5-10 секунд работы
		// CPU-intensive операция
		sum := 0
		for j := 0; j < 10000; j++ {
			sum += j * j
		}
		time.Sleep(100 * time.Millisecond)
	}

	s.mu.Lock()
	s.status.Initialized = true
	s.status.Uptime = time.Since(s.status.StartTime)
	s.mu.Unlock()

	log.Printf("Инициализация завершена за %.2f секунд", s.status.Uptime.Seconds())
}

// IsInitialized проверяет, завершена ли инициализация
func (s *AppService) IsInitialized() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.status.Initialized
}

// GetStatus возвращает текущий статус приложения
func (s *AppService) GetStatus() *models.AppStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Обновляем uptime
	status := *s.status
	status.Uptime = time.Since(s.status.StartTime)

	return &status
}

// GetMemoryUsage возвращает использование памяти в MB
func (s *AppService) GetMemoryUsage() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return float64(m.Alloc) / 1024 / 1024
}

// GetCPUUsage возвращает использование CPU (имитация)
func (s *AppService) GetCPUUsage() float64 {
	// Имитируем стабильное потребление CPU ~0.1
	return 0.1
}

// ProcessLoad имитирует обработку нагрузки
func (s *AppService) ProcessLoad() error {
	if !s.IsInitialized() {
		return ErrAppNotReady
	}

	// Имитируем обработку запроса
	time.Sleep(100 * time.Millisecond)
	return nil
}
