package services

import (
	proto "server1/proto"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	consul "github.com/micro/go-plugins/registry/consul/v2"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8521"), // consul的地址
	)
}

// 创建所有的rpc服务
func Run() {
	// 注册goods 服务
	service := micro.NewService(
		micro.Name("goods"),
		micro.Address(":60441"),
		micro.Registry(consulReg), // 将consul服务注册加入
	)

	service.Init()

	//注册服务 参数1为micro服务
	proto.RegisterGoodsServiceHandler(service.Server(), new(RequestGoods))

	service.Run()
}
