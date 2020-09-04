package util

import (
	"github.com/gin-gonic/gin"
	"golangPro/golang-mallapi/pkg/status"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	if msg == "" {
		msg = status.GetMsg(code)
	}
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}
