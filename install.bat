@echo off
setlocal
cd src
if exist go.mod goto build
go mod init fdsa.ltd/ng
goto build

:build
REM set OLDGOPATH=%GOPATH%
REM set GOPATH=%~dp0


gofmt -w ./

if "%1" == "1" (
    set GOOS=linux
    set GOARCH=amd64
    go build -o ../bin/gn fdsa.ltd/ng 
    goto end
)
set GOOS=windows
set GOARCH=amd64
go build -o ../bin/gn.exe fdsa.ltd/ng 
:end
cd ..
echo finished