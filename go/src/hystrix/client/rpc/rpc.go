package rpc

import (
	proto "client/proto"
	"context"
	"fmt"

	hh "client/wrapper/hystrix"
	"client/wrapper/login"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	consul "github.com/micro/go-plugins/registry/consul/v2"
)

var goodsCli proto.GoodsService

var consulReg registry.Registry

// 所有rpc调用都需要基于这个区调用，先初始化
func init() {

	//服务端注册和客户端注册一致
	consulReg = consul.NewRegistry(
		registry.Addrs("127.0.0.1:8521"), // consul的地址
	)

	microWra := micro.NewService(
		micro.Name("goods"), // 定义服务名
		micro.WrapClient(hh.NewClientWrapper()),
		micro.WrapClient(login.NewClientWrapper()), // 定义中间件
	)

	// 使用中间件重新生成rpc的客户端
	goodsCli = proto.NewGoodsService("goods", microWra.Client())

}

// rpc调用GetGoodDetails
func GetGoodsDetails(goodsId int) (*proto.ResponseGoods, error) {
	fmt.Println("rpc  goodsid ", goodsId)
	res, err := goodsCli.GetGoodsDetails(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})

	return res, err
}
