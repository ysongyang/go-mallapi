package discount

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/pkg/util"
	"strconv"
)

type levelParams struct {
	UserId int `json:"user_id"`
}

// @title 获取推广大使等级
func GetLevels(context *gin.Context) {
	utilGin := util.Gin{Ctx: context}

	UserId, _ := strconv.Atoi(context.Query("UserId"))

	users := &levelParams{
		UserId: UserId,
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utilGin.Response(1, err.Error(), "")
	}
	err = validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utilGin.Response(1, err.Translate(trans), "")
			return
		}
	}

	modelLevel := models.MallDiscountUserLevel{}
	res := modelLevel.GetLevels(UserId)
	utilGin.Response(0, "success", res)
}
