package v1

import (
	"encoding/json"
	"fmt"
	"strconv"

	"client/rpc"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func GetGoodsWra(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id")) // 将接收到的goods_id 转换为int

	var goodsInfo interface{} // 在函数内部和外部都要使用，定义成全局的
	goodsInfo, err := rpc.GetGoodsDetails(goodsId)
	fmt.Println("goodsInfo", goodsInfo)
	goodsInfo = "redis"
	// 连接
	rc, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer rc.Close()

	data, _ := json.Marshal(goodsInfo)

	// 写入数据
	rc.Do("Set", "goods", data)

	// 这里可能是hystrix抛出异常，或者是rpc调用后返回的异常
	if err != nil {
		c.JSON(200, err.Error())
		panic(1)
	}

	c.JSON(200, goodsInfo)
}
