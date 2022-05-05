package api

import (
	apiV1 "client/app/api/v1"
	"client/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	routes.RegisterRoute(Routes)
}

func Routes(g *gin.Engine) {
	//g.GET("getgoods",GetGoods)
	goods := g.Group("goods")
	{
		goods.GET("/", apiV1.GetGoods)
	}

	goods.GET("/wr", apiV1.Wredis)
}
