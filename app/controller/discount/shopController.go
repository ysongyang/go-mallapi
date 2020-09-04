package discount

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"golangPro/golang-mallapi/app/dao"
	"golangPro/golang-mallapi/pkg/status"
	"golangPro/golang-mallapi/pkg/util"
	"strconv"
)

// @title    获取推荐商家
func GetShopList(context *gin.Context) {

	utilGin := util.Gin{Ctx: context}

	Lon, _ := strconv.ParseFloat(context.Query("Lon"), 64)
	Lat, _ := strconv.ParseFloat(context.Query("Lat"), 64)
	Name := context.Query("Name")
	TypeId, _ := strconv.Atoi(context.Query("TypeId"))

	shopList := &dao.ListParams{
		Lon:    Lon,
		Lat:    Lat,
		Name:   Name,
		TypeId: TypeId,
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
	}
	err = validate.Struct(shopList)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utilGin.Response(status.ERROR, err.Translate(trans), nil)
			return
		}
	}

	// 查询数据库
	res, err := shopList.List()
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
		return
	}

	utilGin.Response(status.SUCCESS, "success", res)
}

// @title 获取商家详情
func GetShopDetail(context *gin.Context) {
	utilGin := util.Gin{Ctx: context}
	ShopId, _ := strconv.Atoi(context.Query("ShopId"))

	shopList := &dao.DetailParams{
		ShopId: ShopId,
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
	}
	err = validate.Struct(shopList)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utilGin.Response(status.ERROR, err.Translate(trans), nil)
			return
		}
	}
	// 查询数据库
	res := shopList.Detail()
	if res.Id == 0 {
		utilGin.Response(status.ERROR, "店铺信息不存在或正在审核中.", nil)
		return
	}
	utilGin.Response(status.SUCCESS, "success", res)
}
