package v1

import (
	"github.com/gin-gonic/gin"
)

func InsertNewUser(context *gin.Context) {
	//utilGin := response.Gin{Ctx: context}
	/*users := &model.Users{
		Username: context.PostForm("Username"),
		Password: context.PostForm("Password"),
		Email:    context.PostForm("Email"),
		Mobile:   context.PostForm("Mobile"),
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

	// 写入数据库
	res, err := users.InsertUser()
	if err != nil {
		utilGin.Response(1, err.Error(), "")
		return
	}
	utilGin.Response(0, "success", res)*/
}
