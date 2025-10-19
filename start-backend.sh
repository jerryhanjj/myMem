#!/bin/bash

echo "========================================="
echo "   启动 MyMem 后端服务"
echo "========================================="

cd "$(dirname "$0")/backend"

echo "正在检查 Go 环境..."
go version

echo ""
echo "正在安装依赖..."
go mod tidy

echo ""
echo "正在启动后端服务器 (http://localhost:8080)..."
echo "按 Ctrl+C 停止服务器"
echo ""

go run main.go
