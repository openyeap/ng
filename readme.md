# doc 系统



## go使用方法

Golang支持交叉编译，也就是说你在32位平台的机器上开发，可以编译生成64位平台上的可执行程序。

交叉编译依赖下面几个环境变量：



$GOARCH    目标平台（编译后的目标平台）的处理器架构（386、amd64、arm）
$GOOS          目标平台（编译后的目标平台）的操作系统（darwin、freebsd、linux、windows）



```cmd
# build for linux
set GOOS=linux
set GOARCH=amd64
go build app.go

# build for windows
set GOOS=windows
set GOARCH=amd64
go build app.go


```

