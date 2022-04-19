package main

import (
	_ "mirco_learn/app"
	"mirco_learn/routes"
)

// import引入的时候，线引入app包，就加载app.go app.go中import了api，会调用api/route.go的init方法，然后就会把good的路由注册进来
func main() {
	//初始化路由
	gin := routes.InitRoutes()

	gin.Run()
}
