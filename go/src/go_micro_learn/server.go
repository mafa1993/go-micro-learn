package main

import (
	"fmt"
	pb "go_micro_learn/proto"

	"context"

	micro "github.com/micro/go-micro/v2"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.HelloRequest, res *pb.HelloResponse) error {
	fmt.Println("server")
	res.Greeting = "接收到的值为" + req.Name

	return nil

}

func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name("greeter.server"), // 服务名为greeter.server
		micro.Address("0.0.0.0:9630"),
	)

	// 初始化服务
	service.Init()

	//注册服务 参数1为micro服务
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	service.Run()
}
