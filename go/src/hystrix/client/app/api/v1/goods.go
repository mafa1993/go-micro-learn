package v1

import (
	"fmt"
	"strconv"

	"client/rpc"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id")) // 将接收到的goods_id 转换为int

	// rpc 请求超过一秒就熔断

	// 定义hystrix配置,
	configHy := hystrix.CommandConfig{
		Timeout: 1000, // 超时时间  默认1000毫秒
	}

	// 对Chystrix的指令进行配置
	hystrix.ConfigureCommand("goodsHy", configHy)

	var goodsInfo interface{} // 在函数内部和外部都要使用，定义成全局的
	// 使用前面定义的命令，第二个参数为这个命令的具体做什么操作，第三个参数为服务降级操作
	err := hystrix.Do("goodsHy", func() error {
		// 在发起rpc请求之前，使用hystrix，定义服务熔断和降级
		goodsInfo, err := rpc.GetGoodsDetails(goodsId)
		fmt.Println("goodsInfo", goodsInfo)
		return err
	}, func(e error) error {
		// 降级方法，降级方法不建议使用复杂业务
		// 一般读文件缓存或者是读取缓存
		fmt.Println("降级")
		goodsInfo = demotion()
		return nil
	})

	// 这里可能是hystrix抛出异常，或者是rpc调用后返回的异常
	if err != nil {
		c.JSON(200, err.Error())
		panic(1)
	}

	c.JSON(200, goodsInfo)
}
