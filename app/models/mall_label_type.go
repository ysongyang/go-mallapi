package models

import (
	"golangPro/golang-mallapi/app/models/request"
	"gorm.io/gorm"
)

// 首页活动模块区域分类表模型
type LabelType struct {
	Model
	Id           int    `gorm:"primary_key" json:"id"`
	MallId       int    `json:"mall_id"`
	Name         string `json:"name"`
	Images       string `json:"images"` // 背景图
	Status       string `json:"-"`      // normal开启 hidden关闭
	Weigh        int    `json:"weigh"`  // 权重， 越大越靠前
	CreateTime   int    `json:"-"`
	UpdateTime   int    `json:"-"`
	MasterBanner string `json:"master_banner"` // 内页banner
	//GoodsListResp *GoodsListResp `json:"goods_list"`
	GoodsList []LitestoreGoods `gorm:"ForeignKey:LabelId" json:"goods_list"` //查询当前分类下的商品集合
}

type GoodsListResp struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

func (LabelType) TableName() string {
	return "label_type"
}

// 获取标签下的商品列表
func (labelModel *LabelType) GetAll(params *request.IndexParams) (labelTypes []*LabelType) {
	err := db.Debug().Model(&labelTypes).
		Preload("GoodsList", func(query *gorm.DB) *gorm.DB {
			return query.Order("goods_id desc")
		}).
		Preload("GoodsList.Category").
		Where("mall_id = ? and status = ?", params.MallId, "normal").
		Order("weigh desc").
		Find(&labelTypes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	/*limit := params.LabelGoodsLimit
	  page := params.LabelGoodsPage
	  if limit == 0 {
	  	limit = 9
	  }
	  if page == 0 {
	  	page = 1
	  }
	  for k, v := range labelTypes {
	  	goodsParams := &request.GoodsListParams{
	  		MallId:     params.MallId,
	  		Limit:      limit,
	  		Page:       page,
	  		LabelId:    v.Id,
	  		UserId:     0,
	  		CategoryId: 0,
	  	}
	  	list, count, _ := GetGoodsList(goodsParams)
	  	goodsListRsp := &GoodsListResp{
	  		List:  list,
	  		Total: count,
	  	}
	  	labelTypes[k].GoodsListResp = goodsListRsp

	  }*/
	return labelTypes
}

func GetLabels() []LabelType {
	var labelTypes []LabelType
	err := db.Debug().Model(&LabelType{}).
		Preload("GoodsList", func(query *gorm.DB) *gorm.DB {
			return query.Order("goods_id desc")
		}).
		Preload("GoodsList.Category").
		Where("mall_id = ? and status = ?", 18, "normal").
		Order("weigh desc").
		Find(&labelTypes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return labelTypes
}
