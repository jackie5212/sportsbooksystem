#!/bin/bash

# 网球预定系统 - 快速启动脚本

echo "======================================"
echo "  网球预定系统 - 快速启动"
echo "======================================"
echo ""

# 检查MySQL是否运行
echo "1. 检查MySQL连接..."
if ! command -v mysql &> /dev/null; then
    echo "错误: 未找到mysql命令,请确保MySQL已安装"
    exit 1
fi

# 导入数据库
echo "2. 导入数据库..."
if [ -f "database/schema.sql" ]; then
    mysql -u root -p < database/schema.sql
    if [ $? -eq 0 ]; then
        echo "✓ 数据库导入成功"
    else
        echo "✗ 数据库导入失败"
        exit 1
    fi
else
    echo "错误: 找不到database/schema.sql文件"
    exit 1
fi

echo ""

# 配置后端
echo "3. 配置后端..."
cd server

if [ ! -f "config.yaml" ]; then
    echo "创建配置文件..."
    cp config.yaml.example config.yaml
    echo "✓ 配置文件已创建,请编辑config.yaml修改数据库密码"
    echo ""
    read -p "按回车键继续..."
fi

# 安装依赖
echo "4. 安装Go依赖..."
go mod download
if [ $? -eq 0 ]; then
    echo "✓ 依赖安装成功"
else
    echo "✗ 依赖安装失败"
    exit 1
fi

echo ""
echo "======================================"
echo "  准备启动后端服务..."
echo "======================================"
echo ""
echo "提示:"
echo "1. 确保MySQL正在运行"
echo "2. 确保config.yaml配置正确"
echo "3. 按 Ctrl+C 停止服务"
echo ""

# 启动服务
go run cmd/main.go
