#!/bin/bash
echo "TOTAL Protocol: Launching Sentinel Sandbox..."
echo "------------------------------------------"

# Компиляция симулятора
g++ simulation/sentinel_mock.cpp -o simulation/sentinel_sim

# Запуск в фоновом режиме
./simulation/sentinel_sim &

# Запуск Sentinel Lite (Go-слой)
echo "Starting Sentinel Lite (Execution Layer)..."
cd sentinel-lite && go run cmd/main.go
