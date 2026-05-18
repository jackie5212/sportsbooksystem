@echo off
set GOROOT=C:\Program Files\Go
set PATH=C:\Program Files\Go\bin;%PATH%
set GO111MODULE=on
cd /d %~dp0

REM 删除可能冲突的 go.work 文件
if exist go.work del go.work
if exist go.work.sum del go.work.sum

REM 清理并重新整理依赖
go mod tidy

REM 启动服务
go run cmd\main.go