@echo off
REM 清除所有Go相关环境变量
set GOROOT=
set GOPATH=
set GO111MODULE=
set GOTOOLCHAIN=
set GOFLAGS=

REM 设置正确的路径
set GOROOT=C:\Program Files\Go
set PATH=C:\Program Files\Go\bin;%PATH%
set GO111MODULE=on

cd /d %~dp0

echo 正在清理缓存...
go clean -cache -modcache -i -r

echo.
echo 正在重新下载依赖...
go mod tidy

echo.
echo 正在启动服务...
go run cmd\main.go

pause
