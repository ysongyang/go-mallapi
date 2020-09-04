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

// @title 获取订单列表
func GetOrderList(context *gin.Context) {
	utilGin := util.Gin{Ctx: context}

	UserId, _ := strconv.Atoi(context.Query("UserId"))
	Page, _ := strconv.Atoi(context.Query("Page"))
	Limit, _ := strconv.Atoi(context.Query("Limit"))
	OrderStatus, _ := strconv.Atoi(context.Query("OrderStatus"))

	orderListParasm := &dao.OrderListParams{
		UserId:      UserId,
		Page:        Page,
		Limit:       Limit,
		OrderStatus: OrderStatus,
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
	}
	err = validate.Struct(orderListParasm)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			utilGin.Response(status.ERROR, err.Translate(trans), nil)
			return
		}
	}

	// 查询数据库
	res, err := orderListParasm.GetOrderList()
	if err != nil {
		utilGin.Response(status.ERROR, err.Error(), nil)
		return
	}

	utilGin.Response(status.SUCCESS, "success", res)
}
