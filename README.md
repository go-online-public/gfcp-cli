# gfcp-cli 工具

## 1. 安装

执行以下指令安装 gfcp-cli 至 $GOPATH/bin

```
go install github.com/go-online-public/gfcp-cli@latest
```

## 2. 功能概览

### 2.1 应用模板介绍

```
gfcp-cli newApp .
```

在当前目录下创建应用模板: 

```
.
├── Makefile
├── api
│   ├── api.pb.go
│   ├── api.proto
│   └── api_triple.pb.go
├── cmd
│   └── app.go
├── conf
│   └── dubbogo.yaml
├── go.mod
└── pkg
    └── service
        └── service.go
```

### 2.2 应用模板介绍

生成项目包括几个目录：

- api：放置接口文件：proto文件和生成的.pb.go文件
- cmd：程序入口
- conf：框架配置
- pkg/service：RPC 服务实现
- Makefile: 提供基础脚本
  - make proto-gen 重新编译pb
  - make tidy
  - make test
  - ...

