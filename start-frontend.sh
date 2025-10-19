#!/bin/bash

echo "========================================="
echo "   启动 MyMem 前端服务"
echo "========================================="

cd "$(dirname "$0")/frontend"

echo "正在检查 Node 环境..."
node --version
npm --version

echo ""
echo "正在安装依赖..."
npm install

echo ""
echo "正在启动前端开发服务器 (http://localhost:5173)..."
echo "按 Ctrl+C 停止服务器"
echo ""

npm run dev
