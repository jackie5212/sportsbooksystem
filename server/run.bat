@echo off
set GO111MODULE=on
set GOPATH=
cd /d %~dp0
go run cmd\main.go
