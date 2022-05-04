package routes

import (
	"client/middleware"

	"github.com/gin-gonic/gin"
)

type Router func(*gin.Engine)

// 使用一个全局变量， 其他部分在注册路由前，使用Register方法，把路由注册进来，然后再调用initRoute 进行路由加载
var routers = []Router{}

func InitRoutes() *gin.Engine {
	r := gin.Default()

	middleware.InitMiddleware(r)

	for _, route := range routers {
		route(r)
	}

	return r
}

// 注册route
func RegisterRoute(routes ...Router) {
	routers = append(routers, routes...)
}
