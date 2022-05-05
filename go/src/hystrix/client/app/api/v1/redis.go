package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var key string = "goods"

//使用go连接reedis，实现数据读写
func Wredis(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id")) // 将接收到的goods_id 转换为int

	//	goodsInfo, err := rpc.GetGoodsDetails(goodsId)

	//连接redis
	rc, _ := redis.Dial("tcp", "39.106.6.99:6379")
	defer rc.Close()

	// var data interface{}
	// if err != nil {
	// 	data, _ = json.Marshal(goodsInfo)
	// } else {
	// 	data, _ = json.Marshal(err)
	// }

	//写入数据
	rc.Do("set", key, goodsId)

	// 这里可能是hystrix抛出异常，或者是rpc调用后返回的异常
	// if err != nil {
	// 	c.JSON(200, err.Error())
	// }

	c.JSON(200, "1")
}
