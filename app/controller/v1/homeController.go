package v1

import (
	"github.com/gin-gonic/gin"
	"golangPro/go-mallapi/app/models"
	"golangPro/go-mallapi/pkg/status"
	"golangPro/go-mallapi/pkg/util"
)

func GetDemo(context *gin.Context) {
	utilGin := util.Gin{Ctx: context}
	admin := models.Admin{}
	rest := admin.GetAdmins()
	utilGin.Response(status.SUCCESS, "success", rest)
}
