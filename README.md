# Kubernetes Web Application

Веб-приложение на Go с оптимизированным Kubernetes deployment для мультизонального кластера.

## Архитектура решения

### Требования
- Мультизональный кластер (3 зоны, 5 нод)
- Время инициализации: 5-10 секунд
- Пиковая нагрузка: 4 пода
- CPU: высокое потребление при старте, затем ~0.1 CPU
- Memory: стабильно ~128M
- Дневной цикл нагрузки (пик днем, минимум ночью)
- Максимальная отказоустойчивость + минимальное потребление ресурсов

### Принятые решения

#### 1. Минимальное потребление ресурсов
- **Начальные реплики**: 1 (экономия ресурсов ночью)
- **HPA**: автоматическое масштабирование 1-4 пода
- **Ресурсы**: requests 64Mi/50m, limits 200Mi/200m
- **Graceful shutdown**: 30 секунд для корректного завершения

#### 2. Максимальная отказоустойчивость
- **Pod Anti-Affinity**: распределение по зонам
- **Pod Disruption Budget**: минимум 1 под доступен
- **Rolling Update**: maxUnavailable=1, maxSurge=1
- **Health Checks**: startup, readiness, liveness probes
- **Network Policy**: ограничение сетевого трафика

#### 3. Обработка долгой инициализации
- **Startup Probe**: до 60 секунд на инициализацию
- **Readiness Probe**: проверка готовности к работе
- **Liveness Probe**: проверка работоспособности
- **Graceful handling**: 503 статус до завершения инициализации

#### 4. Дневной цикл нагрузки
- **HPA с поведением**: агрессивное масштабирование вверх, консервативное вниз
- **Метрики**: CPU (70%) и Memory (80%)
- **Стабилизация**: 1 мин вверх, 5 мин вниз

## Структура проекта

`
k8s-app/
├── app/
│   └── main.go          # Go веб-приложение
├── k8s/
│   ├── namespace.yaml      # Namespace
│   ├── configmap.yaml      # Конфигурация
│   ├── deployment.yaml     # Deployment
│   ├── service.yaml        # Service
│   ├── hpa.yaml           # Horizontal Pod Autoscaler
│   ├── pdb.yaml           # Pod Disruption Budget
│   ├── networkpolicy.yaml # Network Policy
│   └── ingress.yaml       # Ingress
├── Dockerfile           # Multi-stage build
├── go.mod              # Go модули
└── README.md           # Документация
`

## Развертывание

### 1. Сборка образа
`ash
docker build -t k8s-app:latest .
`

### 2. Развертывание в Kubernetes
`ash
# Применить все манифесты
kubectl apply -f k8s/

# Проверить статус
kubectl get all -n k8s-app
kubectl get hpa -n k8s-app
kubectl get pdb -n k8s-app
`

### 3. Тестирование
`ash
# Port forward для локального доступа
kubectl port-forward -n k8s-app svc/k8s-app-service 8080:80

# Тестирование endpoints
curl http://localhost:8080/
curl http://localhost:8080/health
curl http://localhost:8080/ready
curl http://localhost:8080/metrics
curl http://localhost:8080/load
`

## Мониторинг

### Метрики приложения
- /metrics - метрики приложения
- /health - health check
- /ready - readiness check

### Kubernetes метрики
`ash
# Статус подов
kubectl get pods -n k8s-app -o wide

# Логи приложения
kubectl logs -n k8s-app -l app.kubernetes.io/name=k8s-app

# HPA статус
kubectl describe hpa -n k8s-app k8s-app-hpa

# События
kubectl get events -n k8s-app --sort-by='.lastTimestamp'
`

## Особенности реализации

### Go приложение
- Имитация долгой инициализации (5-10 сек)
- Высокое потребление CPU при старте
- Стабильное потребление памяти
- Graceful handling запросов до инициализации
- Метрики для мониторинга

### Kubernetes манифесты
- **Безопасность**: non-root пользователь, read-only filesystem, dropped capabilities
- **Отказоустойчивость**: anti-affinity, PDB, rolling updates
- **Экономия ресурсов**: минимальные requests, HPA
- **Сетевая безопасность**: NetworkPolicy с ограничениями
- **Мониторинг**: comprehensive health checks

### Оптимизации
- Multi-stage Docker build для минимального размера образа
- Правильные resource requests/limits
- Startup probe для приложений с долгой инициализацией
- HPA с настроенным поведением для дневного цикла
- Pod Anti-Affinity для распределения по зонам

## Масштабирование

HPA автоматически управляет количеством подов:
- **Минимум**: 1 под (ночью)
- **Максимум**: 4 пода (пиковая нагрузка)
- **Метрики**: CPU 70%, Memory 80%
- **Поведение**: быстрое масштабирование вверх, медленное вниз

## Отказоустойчивость

- **Pod Disruption Budget**: минимум 1 под доступен
- **Anti-Affinity**: распределение по зонам
- **Rolling Updates**: без простоев
- **Health Checks**: автоматическое восстановление
- **Network Policy**: изоляция трафика
