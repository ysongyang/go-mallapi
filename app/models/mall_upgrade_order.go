package models

// 升级VIP订单表[舞星商城升级VIP，士多店铺升级VIP]
type MallUpgradeOrder struct {
	Id            int       `json:"id gorm:"primary_key;auto_increment;"` // 主键ID
	OrderNo       string    `json:"order_no"`                             // 订单ID
	OrderType     string    `json:"order_type"`                           // danceStar 舞星商城VIP升级  shopVip店铺VIP升级 spread会员升级推广大使等级
	MallId        int       `json:"mall_id"`
	ShopId        int       `json:"shop_id"` // 店铺id
	UserId        int       `json:"user_id"` // 会员ID
	LevelId       int       `json:"-"`
	Amount        float32   `json:"amount"`    // 订单金额
	Payamount     float32   `json:"payamount"` // 支付金额
	Paytype       string    `json:"paytype"`   // 支付类型
	Paytime       int       `json:"paytime"`   // 支付时间
	TransactionId string    `json:"-"`
	Ip            string    `json:"ip"` // IP地址
	Type          string    `json:"type"`
	Memo          string    `json:"memo"`   // 备注
	Status        string    `json:"status"` // 状态 created=新建，paid=已支付，expired=过期
	IsRefund      int       `json:"-"`      // 1正常 -1退款
	RefundTime    int       `json:"-"`      // 退款时间
	NotifyResult  string    `json:"-"`
	Createtime    int       `json:"createtime"` // 添加时间
	Updatetime    int       `json:"updatetime"` // 更新时间
	BaseModel     BaseModel `json:"-"`
}

func (MallUpgradeOrder) TableName() string {
	return "upgrade_order"
}

// @title 获取第三方订单表
// @params UserId 用户id
func (model *MallUpgradeOrder) GetUpgradeOrder(UserId int) *MallUpgradeOrder {
	db := model.BaseModel.GetDbConnect("")
	db = db.Where("order_type = ? ", "spread")
	db = db.Where("user_id = ? ", UserId)
	db = db.Where("status = ? ", "paid")
	db = db.Order("paytime desc")
	db.First(&model)

	return model
}
