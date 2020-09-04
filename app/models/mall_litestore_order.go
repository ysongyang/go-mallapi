package models

// 商品订单表
type LitestoreOrder struct {
	Id                  int             `json:"id" gorm:"primary_key"` // 订单ID
	CompanyId           int             `json:"-"`                     // 公司或个人id
	OrderType           string          `json:"order_type"`            // mall 商城订单  shop 店铺订单 discount 优惠联盟
	OrderNo             string          `json:"order_no"`              // 订单编号
	OldOrderNo          string          `json:"-"`                     // 旧订单id
	RecomMallId         int             `json:"-"`                     // 推荐商城id
	MallId              int             `json:"mall_id"`               // 商城Id
	ShopId              int             `json:"shop_id"`               // 店铺id
	UserId              int             `json:"user_id"`               // 用户ID
	DistributionType    int             `json:"-"`                     // 1自提 2店铺配送 3达达配送 4 快递配送
	DistributionPointId int             `json:"-"`                     // 配送点id， 商城id
	RecomUid            int             `json:"-"`                     // 成交推荐人uid
	OrderStatus         string          `json:"order_status"`          // 订单状态:10=待支付,20=已支付,25=已接单,30=配送中,40=待收货,50=已收货,60=已完成,70=退款中,80=退款完成,90=支付失败,100=取消
	IsSettlement        int             `json:"-"`                     // 是否已结算 1结算 0未结算
	TotalPrice          float32         `json:"total_price"`           // 订单总金额
	GoodsPrice          float32         `json:"goods_price"`           // 商品总价
	PayPrice            float32         `json:"pay_price"`             // 订单总支付金额（折扣金额）
	DeliveryAmount      float32         `json:"-"`                     // 配送费
	DeliveryFee         float32         `json:"-"`                     // 实际配送费
	PackFee             float32         `json:"-"`                     // 包装费
	CouponId            int             `json:"-"`                     // 优惠券id
	CouponFee           float32         `json:"coupon_fee"`            // 优惠券金额
	LuckAmount          float32         `json:"luck_amount"`           // 红包减免抵扣金额
	LuckRandomRatio     float32         `json:"-"`                     // 红包减免的系数
	PlatformDed         float32         `json:"platform_ded"`          // 平台补贴
	IntegralDed         int             `json:"integral_ded"`          // 积分抵扣
	MoneyDed            float32         `json:"money_ded"`             // 余额抵扣
	PayTime             int             `json:"pay_time"`              // 支付时间
	PayStatus           string          `json:"pay_status"`            // 支付状态:10=未支付,20=已支付
	ExpressPrice        float32         `json:"-"`                     // 邮费
	ExpressCompany      string          `json:"-"`                     // 快递公司:SF=顺丰速运，BSKD=百世快递，ZTO=中通快递，STO=申通快递，YTO=圆通速递，YD=韵达速递，YZPY=邮政快递，EMS=EMS，HHTT=天天快递，DBL=德邦快递
	ExpressNo           string          `json:"-"`                     // 快递单号
	ExpressTraces       string          `json:"-"`
	Remark              string          `json:"remark"`       // 订单备注
	FreightStatus       string          `json:"-"`            // 发货状态:10=未发货,20=已发货
	FreightTime         int             `json:"-"`            // 发货时间
	ReceiptStatus       string          `json:"-"`            // 收货状态:10=未收货,20=已收货
	PaymentTime         int             `json:"payment_time"` // 发起支付时间
	RefundTime          int             `json:"-"`            // 申请退款时间
	ReceiptTime         int             `json:"-"`            // 收货时间
	ServiceTime         int             `json:"-"`            // 送达时间
	AddressId           int             `json:"address_id"`   // 订单配送地址id
	DistributionTime    string          `json:"-"`            // 送货时间
	TransactionId       string          `json:"-"`            // 微信支付ID
	PayType             string          `json:"-"`
	Type                string          `json:"-"`
	Source              string          `json:"-"` // 来源
	Delivery            int             `json:"-"` // 生鲜订单专用：0 系统默认订单  1 店铺配送 2 自提
	IsRollback          int             `json:"-"` // 0未回滚 1回滚
	QuickPayment        int             `json:"-"` // 0  快捷支付1
	NotifyResult        string          `json:"-"`
	PayIp               string          `json:"-"`          // ip
	IsSplit             int             `json:"-"`          // 是否为拆分订单  0不是， 1是
	IsDelete            int             `json:"-"`          // 是否删除 1正常 -1删除
	Createtime          int             `json:"createtime"` // 生成时间
	Updatetime          int             `json:"-"`
	DbModel             BaseModel       `json:"-"`
	ConsumeDed          float32         `json:"consume_ded"` //平台补贴
	ResultStoreInfo     ResultStoreInfo `json:"store_info"`
}

// 设置表名
func (LitestoreOrder) TableName() string {
	return "litestore_order"
}
