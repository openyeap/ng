@echo off
setlocal
if exist go.mod goto build
go mod init fdsa.ltd/ng
goto build

:build
set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w ./

if "%1" == "1" (
    set GOOS=linux
    set GOARCH=amd64
    go build -o ../bin/ng fdsa.ltd/ng 
    goto end
)
set GOOS=windows
set GOARCH=amd64
go build -o ../bin/ng.exe fdsa.ltd/ng 
:end
set GOPATH=%OLDGOPATH%
echo finished