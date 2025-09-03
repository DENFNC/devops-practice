#!/bin/bash

# Скрипт для развертывания k8s-app в Kubernetes

echo "🚀 Развертывание k8s-app в Kubernetes..."

# Проверяем наличие kubectl
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl не найден. Установите kubectl и настройте доступ к кластеру."
    exit 1
fi

# Проверяем подключение к кластеру
if ! kubectl cluster-info &> /dev/null; then
    echo "❌ Нет подключения к Kubernetes кластеру."
    exit 1
fi

echo "✅ Подключение к кластеру установлено"

# Применяем манифесты в правильном порядке
echo "📦 Создание namespace..."
kubectl apply -f k8s/01-namespace.yaml

echo "⚙️  Создание ConfigMap..."
kubectl apply -f k8s/02-configmap.yaml

echo "🚀 Развертывание приложения..."
kubectl apply -f k8s/03-deployment.yaml

echo "🌐 Создание Service..."
kubectl apply -f k8s/04-service.yaml

echo "📈 Настройка HPA..."
kubectl apply -f k8s/05-hpa.yaml

echo "🛡️  Создание Pod Disruption Budget..."
kubectl apply -f k8s/06-pdb.yaml

echo "🔒 Настройка Network Policy..."
kubectl apply -f k8s/07-networkpolicy.yaml

echo "🌍 Создание Ingress..."
kubectl apply -f k8s/08-ingress.yaml

echo ""
echo "✅ Развертывание завершено!"
echo ""
echo "📊 Статус развертывания:"
kubectl get all -n k8s-app

echo ""
echo "🔍 Для мониторинга используйте:"
echo "  kubectl get pods -n k8s-app -w"
echo "  kubectl logs -n k8s-app -l app.kubernetes.io/name=k8s-app -f"
echo "  kubectl get hpa -n k8s-app"
echo ""
echo "🌐 Для доступа к приложению:"
echo "  kubectl port-forward -n k8s-app svc/k8s-app-service 8080:80"
echo "  curl http://localhost:8080/"
