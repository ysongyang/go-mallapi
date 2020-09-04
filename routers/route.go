package routers

import (
	"github.com/gin-gonic/gin"
	"golangPro/golang-mallapi/app/controller"
	"golangPro/golang-mallapi/app/controller/discount"
	v1 "golangPro/golang-mallapi/app/controller/v1"
	"golangPro/golang-mallapi/app/middleware"
	"golangPro/golang-mallapi/pkg/qrcode"
	"golangPro/golang-mallapi/pkg/upload"
	"net/http"
)

func InitRouter() *gin.Engine {

	// 强制日志颜色化
	gin.ForceConsoleColor()

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	//跨域中间件
	router.Use(middleware.CorsMiddleware())

	router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	router.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//账号登录
	router.POST("/login/account", controller.LoginByAccount)
	//图片上传
	router.POST("/upload/image", controller.UploadImage)

	//多商城V1版本路由
	apiV1 := router.Group("/v1")
	{

		apiV1.GET("home/index", v1.GetIndex)

		// 优惠付
		apiDiscount := apiV1.Group("/discount")
		{
			apiDiscount.GET("/shopList", discount.GetShopList)
			apiDiscount.GET("/shopInfo", discount.GetShopDetail)
			apiDiscount.GET("/level", discount.GetLevels)

			//Token校验中间件
			apiDiscount.Use(middleware.TokenMiddleware())
			{
				apiDiscount.GET("/orderList", discount.GetOrderList)
			}

		}
	}

	return router
}
