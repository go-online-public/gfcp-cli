module dubbo-go-app

go 1.17

require (
	dubbo.apache.org/dubbo-go/v3 v3.0.2-0.20220425022048-4a62aa28e4f9
	github.com/dubbogo/grpc-go v1.42.9
	github.com/dubbogo/triple v1.1.8
	google.golang.org/protobuf v1.28.0
	github.com/LaurenceLiZhixin/dubbogo-extension v0.0.0-20220506071118-45fc9bd5ad5d
)

replace dubbo.apache.org/dubbo-go/v3 => github.com/LaurenceLiZhixin/dubbo-go/v3 v3.0.0-rc1.0.20220506070512-95c18f0940bc
