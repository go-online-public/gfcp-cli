package application

const (
	appGoFile = `package main

import (
	_ "dubbo-go-app/pkg/service"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	_ "github.com/LaurenceLiZhixin/dubbogo-extension/protocol/gfcp-pubsub"
	_ "github.com/LaurenceLiZhixin/dubbogo-extension/registry/gfcp-pubsub"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
`
)

func init() {
	fileMap["appGoFile"] = &fileGenerator{
		path:    "./cmd",
		file:    "app.go",
		context: appGoFile,
	}
}
