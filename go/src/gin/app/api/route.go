package api

import (
	apiV1 "mirco_learn/app/api/v1"
	"mirco_learn/routes"

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
}
