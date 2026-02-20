# Используем официальный образ Go для сборки
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod ./
COPY internal/ ./internal/
COPY cmd/ ./cmd/

# Собираем бинарный файл оркестратора
RUN go build -o orchestrator ./cmd/orchestrator/main.go

# Финальный легкий образ для запуска
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/orchestrator .

# Запуск системы
CMD ["./orchestrator"]
