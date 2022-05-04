package rpc

import (
	"context"
	"fmt"
	proto "client/proto"

	"github.com/micro/go-micro/v2/client"
)

var goodsCli proto.GoodsService

// 所有rpc调用都需要基于这个区调用，先初始化
func init() {
	

}

// rpc调用GetGoodDetails
func GetGoodsDetails(goodsId int) (*proto.ResponseGoods,error) {
	fmt.Println("rpc  goodsid ", goodsId)
	goodsCli = proto.NewGoodsService("goods", client.DefaultClient)
	res, err := goodsCli.GetGoodsDetails(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})

	return res,err
}
