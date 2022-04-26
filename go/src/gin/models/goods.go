package models

import "gorm.io/gorm"

// 商品查询

// 定义orm 字段
type Goods struct {
	gorm.Model
	GoodsId    int `gorm:"primarykey"` // 对应 goods_id字段 如果表中有goodsid 和goods_id两个字段怎么弄  使用标记 标记为primarykey
	ShopId     int
	CategoryId int
	GoodsName  string
}
