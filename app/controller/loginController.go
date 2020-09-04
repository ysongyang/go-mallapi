package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/pkg/status"
	"golangPro/golang-mallapi/pkg/util"
)

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// @title    账号登录控制器
func LoginByAccount(context *gin.Context) {

	utilGin := util.Gin{Ctx: context}

	Username := context.PostForm("Username")
	Password := context.PostForm("Password")

	users := &LoginForm{
		Username: Username,
		Password: Password,
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
	}
	err = validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utilGin.Response(status.ERROR, err.Translate(trans), nil)
			return
		}
	}
	user := &models.MallUser{}

	// 查询数据库
	res, err := user.LoginByAccount(Username, Password)
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
		return
	}
	//生成token
	authToken, err := util.CreateToken(int64(res.Id))
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
		return
	}
	res.AccessToken = *authToken
	utilGin.Response(status.SUCCESS, "success", res)
}
