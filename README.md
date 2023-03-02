# Truss ![Build Status](https://github.com/douchunrong/truss/workflows/Go/badge.svg?branch=master)

## Fork By

https://github.com/metaverse/truss

Truss handles the painful parts of services, freeing you to focus on the
business logic.

![Everything all the time forever](http://i.imgur.com/FtvVeBG.jpg)

## Install

Currently, there is no binary distribution of Truss, you must install from
source.

To install this software, you must:

1. Install protoc 3 or newer. The easiest way is to
   download a release from [github](https://github.com/google/protobuf/releases)
   and add to `$PATH`.
   Otherwise [install from source.](https://github.com/google/protobuf)
1. Install Truss with

   ```shell
   # or git clone https://github.com/douchunrong/truss.git
   go get -u -d github.com/douchunrong/truss
   cd $GOPATH/src/github.com/douchunrong/truss
   make dependencies
   make
   ```
   On Windows, do the following instead:

   ```shell
   go get -u -d github.com/douchunrong/truss
   cd %GOPATH%/src/github.com/douchunrong/truss
   wininstall.bat
   ```

## Usage

Using Truss is easy. You define your service with [gRPC](http://www.grpc.io/)
and [protoc buffers](https://developers.google.com/protocol-buffers/docs/proto3),
and Truss uses that definition to create an entire service. You can even
add [http annotations](
https://github.com/googleapis/googleapis/blob/928a151b2f871b4239b7707e1bb59258df3fe10a/google/api/http.proto#L36)
for HTTP 1.1/JSON transport!

Then you open the `handlers/handlers.go`,
add you business logic, and you're good to go.

Here is an example service definition: [Echo Service](./_example/echo.proto)

Try Truss for yourself on Echo Service to see the service that is generated:

```
truss _example/echo.proto
```

See [USAGE.md](./USAGE.md) and [TUTORIAL.md](./TUTORIAL.md) for more details.

## Developing

See [DEVELOPING.md](./DEVELOPING.md) for details.

## 配置常见错误

### 1. `proto` 文件中添加

```protobuf
import "github.com/douchunrong/truss/deftree/googlethirdparty/annotations.proto";
```

如果出现 `not found github.com/douchunrong/truss/deftree/googlethirdparty/annotations.proto`
检查 `$GO_PATH` 下有没有本项目


### 2. make 报错
错误：
```shell
apple@appledeMacBook-Pro truss % make
go generate github.com/douchunrong/truss/gengokit/template
gengokit/template/gogenerate.go:1: running "go-bindata": exec: "go-bindata": executable file not found in $PATH
make: *** [gobindata] Error 1
```
原因：未找到 go-bindata 可执行文件
解决办法：需要重新导入环境变量,将  `$GTOPATH/bin` 加入到 `$PATH` 中。（我使用的相关环境变量写到了 ~/.bash_profile ，执行 source ~/.bash_profile 即可）

