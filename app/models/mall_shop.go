package models

import "gorm.io/gorm"

// 商城店铺表
type MallShop struct {
	Model
	Id             int     `json:"id"`
	CompanyId      int     `json:"company_id"`    // 公司id
	UserId         int     `json:"user_id"`       // 申请人user_id或绑定人uid
	Name           string  `json:"name"`          // 商城名称
	Contacts       string  `json:"contacts"`      // 联系人
	Phone          string  `json:"phone"`         // 电话
	Address        string  `json:"address"`       // 地址
	Images         string  `json:"-"`             // 店铺图片
	MallCodeUrl    string  `json:"-"`             // 商城二维码
	MallLogo       string  `json:"-"`             // 商城logo
	OpenMallPacket int     `json:"-"`             // 1 开启红包减免 0 關閉
	CategoryType   int     `json:"category_type"` // 分类页模板：0=分类列表，1=分类+商品列表
	FreeShipping   float32 `json:"-"`             // 满多少包邮
	Status         int     `json:"-"`             // 1 开启 0 关闭
	CreateTime     int     `json:"-"`
	UpdateTime     int     `json:"-"`
}

// 设置MallShop的表名为`mall_shop`
func (MallShop) TableName() string {
	return "shop"
}

// @title 根据店铺id获取商城信息
func (model *MallShop) GetShopInfoById(ShopId int) *MallShop {
	err := db.Model(&model).Where("id = ? ", ShopId).First(&model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return model
}
