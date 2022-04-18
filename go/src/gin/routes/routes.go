package routes

import (
	"micro_learn/middleware"

	"github.com/gin-gonic/gin"
)

type Router func(*gin.Engine)

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
func Register(routes ...Router) {
	routers = append(routers, routes...)
}
