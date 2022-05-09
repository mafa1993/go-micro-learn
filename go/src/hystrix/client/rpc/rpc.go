package rpc

import (
	proto "client/proto"
	"context"
	"fmt"

	"client/wrapper/login"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

var goodsCli proto.GoodsService

// 所有rpc调用都需要基于这个区调用，先初始化
func init() {
	microWra := micro.NewService(
		micro.Name("goods.client"),                 // 定义服务名
		micro.WrapClient(login.NewClientWrapper()), // 定义中间件
	)

	// 使用中间件重新生成rpc的客户端
	goodsCli = proto.NewGoodsService("goods",microWra.Client()) 

}

// rpc调用GetGoodDetails
func GetGoodsDetails(goodsId int) (*proto.ResponseGoods, error) {
	fmt.Println("rpc  goodsid ", goodsId)
	res, err := goodsCli.GetGoodsDetails(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})

	return res, err
}
