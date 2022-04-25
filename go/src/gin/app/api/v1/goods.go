package v1

import (
	"mirco_learn/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	goodsId := c.Query("goods_id")

	// 调用服务层获取数据
	goodsInfo := services.GetGoodsDetail(map[string]interface{}{
		"goods_id": goodsId,
	})
	c.JSON(http.StatusOK, goodsInfo)

}
