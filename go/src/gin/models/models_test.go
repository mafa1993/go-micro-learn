package models

import (
	"fmt"
	"testing"
)

func TestGetGoods(t *testing.T) {
	// DB().Debug() 会打印实际执行语句
	/**
	*  where 语句使用  可以传递map[field]interface{}
	*  name=? and age=?
	*  Find(&rlt)  查找并返回数据到rlt
	**/
	where := map[string]interface{}{
		"goods_id": 1,
	}
	var goods Goods
	DB().Debug().Where(where).Find(&goods)
	fmt.Println("goods", goods)
}
