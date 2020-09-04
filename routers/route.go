package routers

import (
	"github.com/gin-gonic/gin"
	v1 "golangPro/go-mallapi/app/controller/v1"
)

func InitRouter() *gin.Engine {

	// 强制日志颜色化
	gin.ForceConsoleColor()

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	//V1版本路由
	apiV1 := router.Group("/v1")
	{

		apiV1.GET("home/index", v1.GetDemo)
	}

	return router
}
