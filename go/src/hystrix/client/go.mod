module client

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo/redis v0.0.1
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	google.golang.org/grpc/examples v0.0.0-20220504170612-799605c22887 // indirect
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
