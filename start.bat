@echo off
chcp 65001 >nul
echo ======================================
echo   网球预定系统 - 快速启动 (Windows)
echo ======================================
echo.

:: 检查MySQL
echo 1. 检查MySQL连接...
where mysql >nul 2>&1
if %errorlevel% neq 0 (
    echo 错误: 未找到mysql命令,请确保MySQL已安装并添加到PATH
    pause
    exit /b 1
)

:: 导入数据库
echo 2. 导入数据库...
if exist database\schema.sql (
    mysql -u root -p < database\schema.sql
    if %errorlevel% equ 0 (
        echo ✓ 数据库导入成功
    ) else (
        echo ✗ 数据库导入失败
        pause
        exit /b 1
    )
) else (
    echo 错误: 找不到database\schema.sql文件
    pause
    exit /b 1
)

echo.

:: 配置后端
echo 3. 配置后端...
cd server

if not exist config.yaml (
    echo 创建配置文件...
    copy config.yaml.example config.yaml
    echo ✓ 配置文件已创建,请编辑config.yaml修改数据库密码
    echo.
    pause
)

:: 安装依赖
echo 4. 安装Go依赖...
call go mod download
if %errorlevel% equ 0 (
    echo ✓ 依赖安装成功
) else (
    echo ✗ 依赖安装失败
    pause
    exit /b 1
)

echo.
echo ======================================
echo   准备启动后端服务...
echo ======================================
echo.
echo 提示:
echo 1. 确保MySQL正在运行
echo 2. 确保config.yaml配置正确
echo 3. 按 Ctrl+C 停止服务
echo.

:: 启动服务
call go run cmd\main.go

pause
