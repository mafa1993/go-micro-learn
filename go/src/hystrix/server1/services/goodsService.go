package services

import (
	"context"
	"encoding/json"
	"fmt"
	"server1/models"
	proto "server1/proto"
)

type RequestGoods struct {
}

func (r *RequestGoods) GetGoodsDetails(ctx context.Context, req *proto.RequestGoods, res *proto.ResponseGoods) error {
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
