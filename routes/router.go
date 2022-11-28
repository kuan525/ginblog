package router

import (
	"ginbolg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	// New也可以,Default多加了两个文件，日志文件等等
	r := gin.Default()

	// 路由分组
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		})
	}

	r.Run()
}
