
# TOTAL Protocol: Sentinel Core v.8.1 Execution Environment
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache gcc musl-dev g++ make

WORKDIR /app
COPY . .

# Собираем основной оркестратор (sentinel-lite)
RUN cd sentinel-lite && go build -o /sentinel_node cmd/main.go

# Собираем симулятор аппаратных инвариантов (C++)
RUN g++ -O3 simulation/main_sim.cpp -o /hardware_sim

# Финальный образ
FROM alpine:latest
RUN apk add --no-cache bash

WORKDIR /root/
COPY --from=builder /sentinel_node .
COPY --from=builder /hardware_sim .
COPY run_demo.sh .

RUN chmod +x run_demo.sh

# При запуске контейнера стартует демонстрация всех уровней защиты
ENTRYPOINT ["./run_demo.sh"]
