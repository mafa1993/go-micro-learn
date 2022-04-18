package main

import (
	_ "mirco_learn/app"
	"mirco_learn/routes"
)

func main() {
	//初始化路由
	gin := routes.InitRoutes()

	gin.Run()
}
