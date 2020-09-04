package dao

import (
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/app/service"
)

// @title 定义获取订单列表的参数结构体
type OrderListParams struct {
	UserId      int `json:"user_id" validate:"required,min=1"`
	Page        int `json:"page" validate:"required,min=1"`  //页码
	Limit       int `json:"limit" validate:"required,min=5"` //分页条数
	OrderStatus int //订单状态
}

type ResultList struct {
	List  []*models.LitestoreOrder `json:"list"`
	Total int                         `json:"total"`
}

// @title 获取订单列表
func (params *OrderListParams) GetOrderList() (*ResultList, error) {
	order := models.LitestoreOrder{}
	db := order.DbModel.GetDbConnect("")

	db = db.Where("is_rollback = ?", 0)
	db = db.Where("is_delete = ?", 1)
	db = db.Where("order_type = ?", "discount")
	db = db.Where("user_id = ?", params.UserId)

	if params.OrderStatus != 0 {
		if params.OrderStatus == 10 {
			//$query->whereIn("order_status", ["10", "90"]);
			db = db.Where("order_status in (?)", []string{"10", "90"})
		}
		if params.OrderStatus == 100 {
			db = db.Where("order_status in (?)", []string{"60", "80"})
		}
	}
	var orders []*models.LitestoreOrder
	res := &ResultList{}

	db = db.Order("createtime desc")
	var count int
	err := db.Find(&orders).Count(&count).Error
	if err != nil {
		return res, err
	}

	if params.Page > 0 && params.Limit > 0 {
		db = db.Limit(params.Limit).Offset((params.Page - 1) * params.Limit)
	}

	if err := db.Find(&orders).Error; err != nil {
		return res, err
	}
	orders = service.MakeOrderList(orders)

	res.List = orders
	res.Total = count

	return res, nil
}
