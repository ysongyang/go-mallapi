package models

import (
	"golangPro/golang-mallapi/app/models/request"
	"log"
)

type LitestoreGoods struct {
	Model
	GoodsId         int               `json:"goods_id" gorm:"primary_key"`
	MallId          int               `json:"mall_id"`     // 商城id
	LabelId         int               `json:"label_id"`    // 关联mall_label_type
	GoodsName       string            `json:"goods_name"`  // 商品名称
	CategoryId      int               `json:"category_id"` // 分类id
	SmallImage      string            `json:"small_image"` // 商品缩略图
	Images          string            `json:"images"`
	SpecType        string            `json:"spec_type"`         // 商品规格:10=单规格,20=多规格
	Describe        string            `json:"describe"`          // 简述
	Content         string            `json:"content"`           // 描述详情
	SalesInitial    int               `json:"sales_initial"`     // 初始销量
	SalesActual     int               `json:"sales_actual"`      // 实际销量
	GoodsSort       int               `json:"goods_sort"`        // 权重
	LimitType       int               `json:"limit_type"`        // 限购类型（0：每单，1：每个id）
	MaxNum          int               `json:"max_num"`           // 限购数量：每人最多可买数量，0表示不限制
	GoodsStatus     string            `json:"goods_status"`      // 商品状态:10=上架,20=下架，30=售罄
	Createtime      int               `json:"-"`                 // 创建时间
	Updatetime      int               `json:"-"`                 // 更新时间
	PlatformRec     int               `json:"-"`                 // 是否平台推荐 0 否 1 是
	ProductId       int               `json:"-"`                 // 产品库id
	DeductStockType string            `json:"deduct_stock_type"` // 库存计算方式:10=下单减库存,20=付款减库存
	Unit            string            `json:"unit"`              // 单位
	StartSellNum    int               `json:"start_sell_num"`    // 起售数量
	Source          string            `json:"-"`                 // 商品来源business商家
	IsPacket        int               `json:"-"`                 // 是否支持红包1：支持，0：不支持
	IsCoupon        int               `json:"-"`                 // 是否支持优惠卷1：支持，0：不支持
	IsFreight       int               `json:"-"`                 // 是否含运费 1含  0不含
	IsDelete        string            `json:"-"`                 // 是否删除:0=未删除,1=已删除
	OldGoodsId      int               `json:"-"`
	ClickRate       int               `json:"click_rate"`                            // 点击率
	DeliveryId      int               `json:"-"`                                     // 运费模板ID
	ActivityId      string            `json:"-"`                                     // 活动id
	ShareRatio      float32           `json:"share_ratio"`                           // 商品分享比例
	Category        LitestoreCategory `gorm:"foreignkey:CategoryID" json:"category"` //分类表
}

func (LitestoreGoods) TableName() string {
	return "litestore_goods"
}

// @title 获取商品列表
func GetGoodsList(params *request.GoodsListParams) (interface{}, int, error) {
	var goods []*LitestoreGoods
	db := db.Model(&goods).Where("goods_status = ? and is_delete = ?", "10", "0")
	if params.MallId > 0 {
		db = db.Where("mall_id = ?", params.MallId)
	}

	if params.LabelId > 0 {
		db = db.Where("label_id = ?", params.LabelId)
	}

	if params.GoodsName != "" {
		db = db.Where("goods_name LIKE ?", "%"+params.GoodsName+"%")
	}

	if params.CategoryId > 0 {
		var categoryModel LitestoreCategory
		categoryIds, _ := categoryModel.GetClassIds(params.MallId, params.CategoryId)
		db = db.Where("category_id In (?)", categoryIds)
	}

	if params.Sort > 0 {
		switch params.Sort {
		case 1:
			db = db.Order("click_rate DESC")
			break
		case 2:
			db = db.Order("(click_rate + sales_actual) desc")
			break
		case 3:
			db = db.Order("goods_price desc")
			break
		case 4:
			db = db.Order("goods_price asd")
			break
		}
	} else {
		if params.Rand > 0 {
			db = db.Order("rand() ")
		} else {
			db = db.Order("goods_sort desc")
		}
	}

	var count int
	db = db.Count(&count)

	if params.Page >= 0 && params.Limit > 0 {
		db = db.Limit(params.Limit).Offset(params.Page)
	}

	if err := db.Preload("Category").Find(&goods).Error; err != nil {
		log.Println(err)
		return nil, 0, nil
	}
	return goods, count, nil
}
