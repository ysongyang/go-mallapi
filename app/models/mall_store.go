package models

import "gorm.io/gorm"

// 店铺记录
type MallStore struct {
	Id                    int       `json:"id" gorm:"primary_key;auto_increment;`
	CompanyId             int       `json:"-"`        // 公司id
	UserId                int       `json:"-"`        // 用户id
	TypeId                int       `json:"-"`        // 店铺类型1生活 2生鲜 3娱乐 4美食
	SupplierId            int       `json:"-"`        // 供应商id
	Name                  string    `json:"name"`     // 店铺名称
	LikeCategoryId        int       `json:"-"`        // 店铺猜你喜欢分类id
	Contacts              string    `json:"contacts"` // 联系人
	RealName              string    `json:"-"`        // 真实姓名 转账时 校验
	Phone                 string    `json:"phone"`    // 联系电话
	Lon                   string    `json:"-"`        // 经度
	Lat                   string    `json:"-"`        // 维度
	Source                string    `json:"-"`
	Geohash               string    `json:"-"`
	Address               string    `json:"address"`              // 地址
	ShopStatus            int       `json:"-"`                    // 1 营业  -1 非营业
	BusinessStartHours    string    `json:"business_start_hours"` // 营业开始时间
	BusinessEndHours      string    `json:"business_end_hours"`   // 营业结束时间
	IsTomorrow            int       `json:"-"`                    // 1:是次日0：是当日
	DistributionFeeQuota  float32   `json:"-"`                    // 配送额度， 减免配送费
	DistributionFee       float32   `json:"-"`                    // 配送费
	DistributionRange     int       `json:"-"`                    // 1 1公里  2 3公里 3 5公里  4 10公里
	Images                string    `json:"images"`
	Remark                string    `json:"remark"`            // 店铺简介
	Status                int       `json:"status"`            // 0 审核中 1 审核通过 -1拒绝通过
	IsPrinter             int       `json:"-"`                 // 0未设置打印机  1 设置打印机
	MachineCode           string    `json:"-"`                 // 设备终端号
	CloseTime             int       `json:"-"`                 // 店铺关闭营业状态的时间记录
	SendTime              int       `json:"-"`                 // 配送时间（分钟）
	Binder                int       `json:"-"`                 // 绑定市场人员id
	BinderName            string    `json:"-"`                 // 市场人员姓名
	ShopCodeUrl           string    `json:"-"`                 // 店铺二维码
	ShopLogo              string    `json:"shop_logo"`         // 店铺logo
	ShopColor             string    `json:"-"`                 // 店铺颜色
	ShopBackground        string    `json:"-"`                 // 背景图
	Cooperator            int       `json:"-"`                 // 合作商6：普通商家10：心有客
	BasicDeliveryFee      int       `json:"-"`                 // 起送费
	OpenLuckMoney         int       `json:"-"`                 // 1开启红包抵扣支付0关闭
	IsDistribution        int       `json:"-"`                 // 1 支持配送 0 不支持配送
	IsDada                int       `json:"-"`                 // 是否支持达达配送 1支持 0不支持
	ShareRatio            float32   `json:"-"`                 // 商品分享相关的分佣比例,0则不开启[收费店铺有效]
	MerchantDedProportion float32   `json:"-"`                 // 商家支付配送费比例
	UserDedProportion     float32   `json:"-"`                 // 用户支付配送费比例
	LuckQuota             float32   `json:"-"`                 // 红包补贴额度（红包抵扣时用）
	SmallPacket           int       `json:"-"`                 // 最小红包使用系数（%）
	MaxPacket             int       `json:"-"`                 // 最大红包使用系数（%）
	IsVip                 int       `json:"-"`                 // 0 普通商户 1 收费商户
	IsOpenDadaShop        int       `json:"-"`                 // 是否开启达达门店 0 : 未开启 1：开启
	ProvinceId            int       `json:"-"`                 //省级id
	CityId                int       `json:"-"`                 //市级id
	RegionId              int       `json:"-"`                 //区域id
	DiscountQrcode        string    `json:"discount_qrcode"`   // 成为优惠联盟商家后的二维码
	IsDiscount            int       `json:"-"`                 // 是否成为优惠联盟商家 1是 0不是
	RechargeAmountCount   float32   `json:"-"`                 // 累计充值金额
	DiscountQuotaTarget   float32   `json:"-"`                 // 商家储值额度（固定）
	DiscountQuotaAmount   float32   `json:"-"`                 // 商家储值额度(用于消费后抵扣)
	ThresholdQuotaAmount  float32   `json:"-"`                 // 储存额度报警阈值
	PlatformDiscount      int       `json:"platform_discount"` // 平台补贴折扣  10
	MerchantDiscount      int       `json:"-"`                 // 商家折扣（商家给平台的折扣）80
	CreateTime            int       `json:"create_time"`
	UpdateTime            int       `json:"-"`
	Distance              float64   `json:"distance"`
	DbModel               BaseModel `json:"-"`
}

// 设置MallStore的表名为`mall_store`
func (MallStore) TableName() string {
	return "store"
}

type ResultStoreInfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`      // 店铺名称
	ShopLogo string `json:"shop_logo"` // 店铺logo
}

// @title 获取店铺信息
func (model *MallStore) GetShopInfo(StoreId int) ResultStoreInfo {
	db := model.DbModel.GetDbConnect("")
	var result ResultStoreInfo
	db.Raw("SELECT id, name, shop_logo FROM mall_store WHERE id = ?", StoreId).Scan(&result)
	return result
}

// @title 根据店铺id获取店铺信息
func (model *MallStore) GetStoreInfoById(StoreId int) *MallStore {
	err := db.Where("id = ? ", StoreId).First(&model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return model
}
