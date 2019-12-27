@echo off
go build -o ./bin/windows/amd64/ng.exe 

start ./bin/windows/amd64/ng.exe 
REM set GOOS=linux
REM set GOARCH=amd64
REM go build -o ./bin/linux/amd64/ng main.go

REM set GOOS=windows
REM set GOARCH=amd64
REM go build -o ./bin/windows/amd64/ng.exe main.go
