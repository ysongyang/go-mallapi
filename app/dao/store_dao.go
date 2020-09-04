package dao

import (
	"github.com/mmcloughlin/geohash"
	"golangPro/golang-mallapi/app/models"
	"golangPro/golang-mallapi/app/service"
	"golangPro/golang-mallapi/pkg/util"
)

// @title 定义获取店铺的参数结构体
type ListParams struct {
	Lon    float64 `json:"lon" validate:"required"` //经度
	Lat    float64 `json:"lat" validate:"required"` //维度
	Name   string  //店铺名称
	TypeId int     //店铺分类
}

// @title 获取店铺列表
func (params *ListParams) List() ([]*models.MallStore, error) {
	store := models.MallStore{}
	db := store.DbModel.GetDbConnect("")

	geohash := geohash.Encode(params.Lat, params.Lon)

	geohash = util.Substr(geohash, 0, 3)

	db = db.Where("is_discount = ?", 1) //过滤只显示优惠付
	db = db.Where("status = ?", 1)      //过滤审核通过

	db = db.Where("geohash LIKE ?", "%"+geohash+"%")

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
		db = db.Or("address LIKE ?", "%"+params.Name+"%")
	}

	if params.TypeId != 0 {
		db = db.Where("type_id = ?", params.TypeId)
	}
	var storeS []*models.MallStore

	db.Find(&storeS)

	//计算距离
	storeS = service.CalLonLatDistance(storeS, params.Lon, params.Lat)
	return storeS, nil
}

// @title 定义店铺详情的参数结构体
type DetailParams struct {
	ShopId int `json:"shop_id" validate:"required,min=1"` //店铺ID
}

// @title 获取店铺详情
func (params *DetailParams) Detail() *models.MallStore {
	store := models.MallStore{}
	db := store.DbModel.GetDbConnect("")
	db = db.Where("is_discount = ?", 1) //过滤只显示优惠付
	db = db.Where("status = ?", 1)      //过滤审核通过
	db = db.Where("ID = ?", params.ShopId)
	db.First(&store)
	return &store
}
