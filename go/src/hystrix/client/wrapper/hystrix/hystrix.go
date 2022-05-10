package hystrix

import (
	"context"
	"encoding/json"
	"fmt"

	proto "client/proto"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/client"
)

type clientWrapper struct {
	client.Client
}

var key string = "goods"

// 中间件方法  第三个参数为返回值  response
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	//这里为中间件要执行的内容
	fmt.Println("hystrix 中间件 执行")
	configHy := hystrix.CommandConfig{
		Timeout: 1000,
	}

	hystrix.ConfigureCommand("goodsWra", configHy)

	return hystrix.Do("goodsWra", func() error {
		err := c.Client.Call(ctx, req, rsp, opts...) // 这里相当于next
		return err
	}, func(e error) error {
		// 降级方法需要给客户端返回内容，返回的内容放在req中
		demotion(rsp)
		return nil
	})

	fmt.Println("hystrix 中间件 结束")

	return nil
}

// 创建中间件对象
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}

}

/**
 *  降级
 */
func demotion(res interface{}) {
	rc, _ := redis.Dial("tcp", "39.106.6.99:6379")
	defer rc.Close()

	// 将redis返回的数据转化昵称byte类型
	data, _ := redis.Bytes(rc.Do("Get", key))

	// response 应该proto中定义的结构体 goodsResponse 需要将数据转换为对应的结构体
	res = res.(*proto.ResponseGoods)

	//将data进行json解码 并存入res
	json.Unmarshal(data, res)

}
