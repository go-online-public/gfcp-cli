package application

const (
	serviceFile = `package service

import (
	"context"
)

import (
	"dubbo-go-app/api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
)

type GreeterServerImpl struct {
	api.UnimplementedGFCPDemoServiceServer
}

func (s *GreeterServerImpl) SayHello(ctx context.Context, in *api.HelloRequest) (*api.User, error) {
	logger.Infof("GFCPProvider get user name = %s\n", in.Name)
	return &api.User{Name: "Hello " + in.Name + "V2", Age: 22}, nil
}

func init() {
	config.SetProviderService(&GreeterServerImpl{})
}
`
)

func init() {
	fileMap["serviceFile"] = &fileGenerator{
		path:    "./pkg/service",
		file:    "service.go",
		context: serviceFile,
	}
}
