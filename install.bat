@echo off
setlocal

if exist go.mod goto build
go mod init fdsa.ltd/ng
goto build

:build
REM set OLDGOPATH=%GOPATH%
REM set GOPATH=%~dp0

REM gofmt -w src

set GOOS=linux
set GOARCH=amd64
go build -o bin/ng fdsa.ltd/ng/src

set GOOS=windows
set GOARCH=amd64
go build -o bin/ng.exe fdsa.ltd/ng/src


:end
REM set GOPATH=%OLDGOPATH%
REM del go.mod
REM del go.sum

echo finished