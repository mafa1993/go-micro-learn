package services

import (
	"context"
	"encoding/json"
	"fmt"
	"server1/models"
	proto "server1/proto"
	"time"
)

type RequestGoods struct {
}

func (r *RequestGoods) GetGoodsDetails(ctx context.Context, req *proto.RequestGoods, res *proto.ResponseGoods) error {
	fmt.Println("60441 开启")
	// 模拟掉线  增加抖动
	time.Sleep(time.Second * 3)

	where := map[string]interface{}{
		"goods_id": req.GoodsId, // goodsId 是在proto中定义的
	}

	var goods models.Goods
	models.DB().Debug().Where(where).Unscoped().Find(&goods)
	json_rlt, err := json.Marshal(goods)

	if err != nil {
		res.Code = 200
		res.Msg = "json转化出错"
		res.Data = []byte{1, 2, 3}
	}

	res.Code = 200
	res.Msg = "成功"
	res.Data = json_rlt

	fmt.Println("data", res)

	return nil
}
