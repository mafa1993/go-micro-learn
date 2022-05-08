package v1

import (
	"encoding/json"
	"strconv"

	"client/rpc"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var key string = "goods"

//使用go连接reedis，实现数据读写
func Wredis(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id")) // 将接收到的goods_id 转换为int

	goodsInfo, _ := rpc.GetGoodsDetails(goodsId)

	//连接redis
	rc, _ := redis.Dial("tcp", "39.106.6.99:6379")
	defer rc.Close()

	goodsInfo.Msg = "redis"
	// var data interface{}
	// if err != nil {
	// 	data, _ = json.Marshal(goodsInfo)
	// } else {
	// 	data, _ = json.Marshal(err)
	// }

	data, _ := json.Marshal(goodsInfo)

	//写入数据
	rc.Do("set", key, data)

	// 这里可能是hystrix抛出异常，或者是rpc调用后返回的异常
	// if err != nil {
	// 	c.JSON(200, err.Error())
	// }

	c.JSON(200, data)
}

// 从redis中读取数据
func Gredis(c *gin.Context) {
	rc, _ := redis.Dial("tcp", "39.106.6.99:6379")
	defer rc.Close()

	data, _ := rc.Do("Get", key)

	c.JSON(200, string(data.([]byte))) // 断言  将interface类型转换成后面括号里的类型，https://blog.csdn.net/raoxiaoya/article/details/115131883
}

/**
 *  降级
 */
func demotion() string {
	rc, _ := redis.Dial("tcp", "39.106.6.99:6379")
	defer rc.Close()

	data, _ := rc.Do("Get", key)

	return string(data.([]byte))
}
