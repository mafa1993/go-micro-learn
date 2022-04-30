package main

import (
	"context"
	"fmt"
	pb "go_micro_learn/proto"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("geeter.client"))

	service.Init()

	// 创建新客户端
	greeter := pb.NewGreeterService("greeter.server", service.Client()) // 定义要访问的服务

	// 调用服务的Hello方法
	res, err := greeter.Hello(context.Background(), &pb.HelloRequest{
		Name: "abc",
	})
	fmt.Println("res", res)
	fmt.Println(err)

}
