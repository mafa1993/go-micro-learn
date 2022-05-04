package main

import (
	_ "client/app"
	_ "client/rpc"
	"client/routes"

	"github.com/micro/go-micro/v2/web" // go micro 中的web包
)

// import引入的时候，线引入app包，就加载app.go app.go中import了api，会调用api/route.go的init方法，然后就会把good的路由注册进来
func main() {
	//初始化路由
	gin := routes.InitRoutes()

	server := web.NewService(
		web.Name("goods.client"),
		web.Handler(gin),
	)

	server.Init()

	gin.Run()
}
