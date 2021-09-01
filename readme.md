# 文档系统

## 在线文档编辑系统

一个简单的在线文档编辑器，充分利用分布式算法，把零散的文章通过版本化管理在云上。再利用计算机的结构管理帮助作者完成创作。

## 使用到的技术

1. git版本控制
1. 分布式算法
1. markdown语法
1. 网页技术

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

