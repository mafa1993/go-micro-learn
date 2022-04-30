package v1

import (
	"fmt"
	"strconv"

	"micro_client/rpc"

	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id")) // 将接收到的goods_id 转换为int

	goodsInfo := rpc.GetGoodsDetails(goodsId)
	fmt.Println("goodsInfo", goodsInfo)

	c.JSON(200, goodsInfo)
}
