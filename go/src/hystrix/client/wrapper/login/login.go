package login

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/client"
)

type clientWrapper struct {
	client.Client
}

// 中间件方法
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	//这里为中间件要执行的内容
	fmt.Println("前置中间件执行")

	err := c.Client.Call(ctx, req, rsp, opts...) // 这里相当于next

	fmt.Println("后置中间件执行")

	return err
}

// 创建中间件对象
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}

}
