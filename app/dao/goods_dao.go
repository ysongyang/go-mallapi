package dao

import (
	"golangPro/golang-mallapi/app/models"
)

type GoodsListField struct {
	GoodsId         int          `json:"goods_id" gorm:"primary_key"`
	MallId          int          `json:"mall_id"`     // 商城id
	GoodsName       string       `json:"goods_name"`  // 商品名称
	CategoryId      int          `json:"category_id"` // 分类id
	SmallImage      string       `json:"small_image"` // 商品缩略图
	Images          string       `json:"images"`
	SpecType        string       `json:"spec_type"`         // 商品规格:10=单规格,20=多规格
	SalesInitial    int          `json:"sales_initial"`     // 初始销量
	SalesActual     int          `json:"sales_actual"`      // 实际销量
	GoodsStatus     string       `json:"goods_status"`      // 商品状态:10=上架,20=下架，30=售罄
	DeductStockType string       `json:"deduct_stock_type"` // 库存计算方式:10=下单减库存,20=付款减库存
	Unit            string       `json:"unit"`              // 单位
	StartSellNum    int          `json:"start_sell_num"`    // 起售数量
	IsFreight       int          `json:"is_freight"`        // 是否含运费 1含  0不含
	ClickRate       int          `json:"click_rate"`        // 点击率
	models          models.Model `json:"-"`
}

