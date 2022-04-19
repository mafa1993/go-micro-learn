package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	c.JSON(http.StatusOK, "成功")

}
