package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"golangPro/golang-mallapi/app/models/request"
	"golangPro/golang-mallapi/app/service"
	"golangPro/golang-mallapi/pkg/status"
	"golangPro/golang-mallapi/pkg/util"
	"strconv"
)

// @title 首页接口
func GetIndex(context *gin.Context) {
	utilGin := util.Gin{Ctx: context}
	MallId, _ := strconv.Atoi(context.Query("MallId"))
	UserId, _ := strconv.Atoi(context.Query("UserId"))
	CategoryId, _ := strconv.Atoi(context.Query("CategoryId"))
	TopCategoryLimit, _ := strconv.Atoi(context.Query("TopCategoryLimit"))
	CategoryLimit, _ := strconv.Atoi(context.Query("CategoryLimit"))
	LabelGoodsPage, _ := strconv.Atoi(context.Query("LabelGoodsPage"))
	LabelGoodsLimit, _ := strconv.Atoi(context.Query("LabelGoodsLimit"))

	indexParams := &request.IndexParams{
		MallId:           MallId,
		UserId:           UserId,
		CategoryId:       CategoryId,
		TopCategoryLimit: TopCategoryLimit,
		CategoryLimit:    CategoryLimit,
		LabelGoodsPage:   LabelGoodsPage,
		LabelGoodsLimit:  LabelGoodsLimit,
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	//注册一个函数，获取struct tag里自定义的label作为字段名
	/*validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name:=fld.Tag.Get("label")
		return name
	})*/

	//验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
	}
	err = validate.Struct(indexParams)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utilGin.Response(status.ERROR, err.Translate(trans), nil)
			return
		}
	}
	rest := service.GetIndex(indexParams)
	utilGin.Response(status.SUCCESS, "success", rest)
}
