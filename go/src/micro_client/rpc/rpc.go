package rpc

import (
	"context"
	"fmt"
	proto "micro_client/proto"

	"github.com/micro/go-micro/v2/client"
)

func init() {

}

func GetGoodsDetails(goodsId int) *proto.ResponseGoods {
	fmt.Println("rpc  goodsid ", goodsId)
	goods := proto.NewGoodsService("goods", client.DefaultClient)

	res, _ := goods.GetGoodsDetails(context.Background(), &proto.RequestGoods{
		GoodsId: int32(goodsId),
	})

	return res
}
