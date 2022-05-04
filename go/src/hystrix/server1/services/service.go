package services

import (
	proto "server1/proto"

	micro "github.com/micro/go-micro/v2"
)

// 创建所有的rpc服务
func Run() {
	// 注册goods 服务
	service := micro.NewService(
		micro.Name("goods"),
	)

	service.Init()

	//注册服务 参数1为micro服务
	proto.RegisterGoodsServiceHandler(service.Server(), new(RequestGoods))

	service.Run()
}
