#!/bin/bash
echo "🚀 Starting TOTAL Protocol: Sentinel Lite Node..."
cd sentinel-lite
go run main.go &
echo "🌐 Opening Frontend Sandbox..."
sleep 2
# Открывает браузер (работает на Mac и большинстве Linux)
open ../frontend/index.html || xdg-open ../frontend/index.html
