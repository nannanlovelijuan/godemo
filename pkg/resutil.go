package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//处理请求和返回的工具类

func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":   200,
			"msg":    "success",
			"data":   data,
			"status": true,
		},
	)
}
