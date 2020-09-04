package middleware

import (
	"github.com/gin-gonic/gin"
	"golangPro/golang-mallapi/pkg/status"
	"golangPro/golang-mallapi/pkg/util"
)

// @title Token中间件
func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		utilGin := util.Gin{Ctx: c}
		Authorization := c.GetHeader("Authorization")
		if Authorization == "" {
			utilGin.Response(status.ERROR_AUTH_CHECK_TOKEN_FAIL, "", nil)
			c.Abort()
			return
		}
		_, err := util.ParseToken(Authorization)
		if err != nil {
			utilGin.Response(status.ERROR_AUTH, "", nil)
			c.Abort()
			return
		}
		c.Next()

		//fmt.Printf("userinfo: %v\n", ret.UserId)

	}
}

// @title 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("CorsMiddleware before next")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("AAccess-Control-Allow-Headers", "DNT,Keep-Alive,User-Agent,Cache-Control,Content-Type,Authorization,client_type")
		//c.Set("name", "test")
		/*
		   c.Next()后就执行真实的路由函数，路由函数执行完成之后继续执行后续的代码
		*/
		c.Next()
		//fmt.Println("CorsMiddleware after next")
	}
}
